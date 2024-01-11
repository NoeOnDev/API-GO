package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/go-pg/pg/v10"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
)

func UserRoutes(router *gin.Engine, db *pg.DB) {
    router.POST("/register", func(c *gin.Context) {
        controllers.Register(c, db)
    })
}