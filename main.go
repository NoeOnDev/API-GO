package main

import (
	"github.com/NoeOnDev/API-GO/controllers"
	"github.com/NoeOnDev/API-GO/database"
	"github.com/NoeOnDev/API-GO/middlewares"
	"github.com/NoeOnDev/API-GO/models"

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

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.POST("/pets", controllers.AddPet)
	protectedRoutes.GET("/pets", controllers.GetAllPets)
	
	r.Run()
}
