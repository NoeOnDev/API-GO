package routes

import (
	"net/http"
	"github.com/NoeAlejandroRodriguezMoto/API-GO/controllers"
)

func ClientRoutes() {
    clientController := controllers.ClientController{}

    http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            clientController.Create(w, r)
        default:
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        }
    })
}