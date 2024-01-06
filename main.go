package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/database"
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

    port := os.Getenv("GO_PORT")
    if port == "" {
        log.Fatal("$GO_PORT must be set")
    }

    log.Printf("Listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}