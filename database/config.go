package database

import (
	"os"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func ConnectDB() {
	pgDB := pg.Connect(&pg.Options{
		Addr:    	os.Getenv("DB_ADDR"),
		User:   	os.Getenv("DB_USER"),
		Password: 	os.Getenv("DB_PASS"),
		Database: 	os.Getenv("DB_NAME"),
	})

	defer pgDB.Close()

	_, err := pgDB.Exec("SELECT 1")
	if err != nil {
		fmt.Println("Error connecting to the database", err)
		return
	}
	fmt.Println("Connection to the database was successful")
}