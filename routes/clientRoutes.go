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
    r.PUT("/clients/:id", func(c *gin.Context) {
        controllers.UpdateClient(c, db)
    })
}

func AuthRoutes(r *gin.Engine, db *sql.DB) {
    r.POST("/login", func(c *gin.Context) {
        controllers.Login(c, db)
    })
}