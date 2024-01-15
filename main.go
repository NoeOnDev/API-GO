package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envLoad()
	databaseLoad()
	serverAplication()
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func databaseLoad() {

}

func serverAplication() {
	r := gin.Default()
	r.Run()
}