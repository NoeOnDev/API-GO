package models

import (
    "time"
)

type Client struct {
	ID		  int       `json:"id"`
    FirstName string    `json:"firstname"`
    LastName  string    `json:"lastname"`
    BirthDate time.Time `json:"birthdate"`
    Phone     string    `json:"phone"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
}