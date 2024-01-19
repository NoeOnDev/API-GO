package main

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envLoad()
	databaseLoad()
	serverAplication()
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func databaseLoad() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Pet{})
}

func serverAplication() {
	r := gin.Default()
	r.Run()
}