package controllers

import (
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/models"
)

type LoginCredentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Login(c *gin.Context, db *sql.DB) {
    var creds LoginCredentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var client models.Client
    err := db.QueryRow("SELECT id, password FROM clients WHERE email = $1", creds.Email).Scan(&client.ID, &client.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(creds.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": client.ID})
}