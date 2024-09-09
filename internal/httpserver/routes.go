package httpserver

import (
	"net/http"

	"github.com/lauragrassig/golang-api/internal/handler"
)

func initRoutes(h *handler.Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /user", h.GetUser)
	mux.HandleFunc("GET /users", h.GetUsers)
	mux.HandleFunc("POST /user", h.CreateUser)
}
