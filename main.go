package main

import (
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/models"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"

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

	publicRoutes := r.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)
	
	r.Run()
}