package controllers

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/models"
)

type ClientController struct{
    DB *sql.DB
}

func (c *ClientController) Create(ctx *gin.Context) {
    var client models.Client

    err := ctx.ShouldBindJSON(&client)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    query := `INSERT INTO client (first_name, last_name, birthdate, email, phone, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    err = c.DB.QueryRow(query, client.First_Name, client.Last_Name, client.Birthdate, client.Email, client.Phone, client.Address).Scan(&client.ID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, client)
}