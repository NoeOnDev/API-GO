package main 

import (
    "os"
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/database"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/routes"
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

    db := database.ConnectDB()

    r := gin.Default()
    routes.ClientRoutes(r, db)
    r.Run(":" + PORT)
}