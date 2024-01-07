package routes

import (
    "net/http"
    "github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
)

func ClientRoutes(c *controllers.ClientController) {
    http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            c.Create(w, r)
        default:
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        }
    })
}