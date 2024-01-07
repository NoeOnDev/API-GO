package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/database"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/routes"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    db := database.ConnectDB()
    defer db.Close()

    clientController := controllers.ClientController{DB: db}
    routes.ClientRoutes(&clientController)

    port := os.Getenv("GO_PORT")
    if port == "" {
        log.Fatal("$GO_PORT must be set")
    }

    log.Printf("Listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}