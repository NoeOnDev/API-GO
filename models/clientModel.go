package models

import "time"

type Client struct {
	ID      	int    `json:"id"`
	First_Name  string `json:"name"`
	Last_Name 	string `json:"last_name"`
	Birthdate 	time.Time `json:"birthdate"`
	Email 		string `json:"email"`
	Phone 		string `json:"phone"`
	Address 	string `json:"address"`
}