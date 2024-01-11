package main

import (
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/database"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    port := os.Getenv("GO_PORT")

    db, err := database.ConnectDB()
    if err != nil {
        log.Fatal("Error connecting to the database", err)
    }
    defer db.Close()

    router := gin.Default()

    routes.UserRoutes(router, db)

    router.Run(":" + port)
}