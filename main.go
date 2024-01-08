package main 

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("GO_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r := gin.Default()
	r.Run(":" + PORT)
}

