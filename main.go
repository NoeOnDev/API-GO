package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GO_PORT")

	database.ConnectDB()
	router := gin.Default()
	router.Run(":" + port)
}