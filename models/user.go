
package models

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Usename string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}