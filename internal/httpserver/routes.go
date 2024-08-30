package httpserver

import (
	"net/http"

	"github.com/lauragrassig/golang-api/internal/handler"
)

func initRoutes(h *handler.Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /user", h.Home)
	mux.HandleFunc("GET /users", h.Home)
	mux.HandleFunc("POST /user", h.Home)
}
