package controllers

import (
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/models"
)

func CreateClient(c *gin.Context, db *sql.DB) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec("INSERT INTO clients (firstname, lastname, birthdate, phone, email, password) VALUES ($1, $2, $3, $4, $5, $6)", client.FirstName, client.LastName, client.BirthDate, client.Phone, client.Email, client.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": client})
}

func GetClients(c *gin.Context, db *sql.DB) {
    rows, err := db.Query("SELECT id, firstname, lastname, birthdate, phone, email, password FROM clients")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var clients []models.Client
    for rows.Next() {
        var client models.Client
        if err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.BirthDate, &client.Phone, &client.Email, &client.Password); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        clients = append(clients, client)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": clients})
}