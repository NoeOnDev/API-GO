package routes

import (
    "net/http"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
    "github.com/gin-gonic/gin"
)

func ClientRoutes(router *gin.Engine, c *controllers.ClientController) {
    router.Any("/client", func(ctx *gin.Context) {
        switch ctx.Request.Method {
        case "POST":
            c.Create(ctx)
        default:
            ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
        }
    })
}