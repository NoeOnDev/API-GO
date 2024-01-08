package routes

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
)

func ClientRoutes(r *gin.Engine, db *sql.DB) {
    r.POST("/clients", func(c *gin.Context) {
        controllers.CreateClient(c, db)
    })
    r.GET("/clients", func(c *gin.Context) {
        controllers.GetClients(c, db)
    })
}