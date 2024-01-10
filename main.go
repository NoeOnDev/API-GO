package main

import (
	"fmt"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	pgDB := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: pass,
		Database: name,
	})

	defer pgDB.Close()
	_, error := pgDB.Exec("SELECT 1")
	if error != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Connected to the database")

	router := gin.Default()
	router.Run(":8080")
}