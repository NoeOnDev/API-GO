package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/go-pg/pg/v10"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/middlewares"
)

func UserRoutes(router *gin.Engine, db *pg.DB) {
    router.POST("/register", middlewares.ValidateUser(), func(c *gin.Context) {
        controllers.Register(c, db)
    })
}