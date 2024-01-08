package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type PingController struct {}

func (PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", PingController{}.Ping)
	r.Run()
}

