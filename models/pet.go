package models

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Category string `gorm:"size:100;not null;unique" json:"category"`
	UserID uint `gorm:"not null" json:"user_id"`
	User User
}

func (pet *Pet) SavePet() (*Pet, error) {
	err := database.Database.Create(&pet).Error
	if err != nil {
		return &Pet{}, err
	}
	return pet, nil
}

func GetAllPetsByUserID(userID uint) ([]Pet, error) {
    var pets []Pet
    err := database.Database.Preload("User").Where("user_id = ?", userID).Find(&pets).Error
    if err != nil {
        return nil, err
    }
    return pets, nil
}