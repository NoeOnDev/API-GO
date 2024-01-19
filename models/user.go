package models

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Pets []Pet
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
 
func (user *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsernameOrEmail(usernameOrEmail string) (User, error) {
    var user User
    err := database.Database.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}

func FindUserWithPetsById(id uint) (User, error) {
    var user User
    err := database.Database.Preload("Pets").Where("ID=?", id).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}