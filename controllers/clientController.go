package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/models"
)

type ClientController struct{
    DB *sql.DB
}

func (c *ClientController) Create(w http.ResponseWriter, r *http.Request) {
    var client models.Client

    err := json.NewDecoder(r.Body).Decode(&client)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `INSERT INTO client (first_name, last_name, birthdate, email, phone, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    err = c.DB.QueryRow(query, client.First_Name, client.Last_Name, client.Birthdate, client.Email, client.Phone, client.Address).Scan(&client.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(client)
}