package database

import (
    "os"
    "fmt"
    "github.com/go-pg/pg/v10"
)

func ConnectDB() (*pg.DB, error) {
    pgDB := pg.Connect(&pg.Options{
        Addr:    	os.Getenv("DB_ADDR"),
        User:   	os.Getenv("DB_USER"),
        Password: 	os.Getenv("DB_PASS"),
        Database: 	os.Getenv("DB_NAME"),
    })

    _, err := pgDB.Exec("SELECT 1")
    if err != nil {
        fmt.Println("Error connecting to the database", err)
        return nil, err
    }
    fmt.Println("Connection to the database was successful")
    return pgDB, nil
}