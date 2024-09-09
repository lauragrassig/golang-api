package httpserver

import (
	"fmt"
	"net/http"

	"github.com/lauragrassig/golang-api/internal/handler"
)

type httpSrv struct{
	handler *handler.Handler
	hostname string
	port string
}

func New(h * handler.Handler, hostname string, port string) *httpSrv{
	return &httpSrv{
		handler: h,
		hostname: hostname,
		port: port,
	}
}

func (h *httpSrv) Start() {
	// addr := h.hostname + ":" + h.port

	mux := http.NewServeMux()
	initRoutes(h.handler, mux)
	addr := fmt.Sprintf("%s:%s", h.hostname, h.port)
	fmt.Printf("Servidor http rodando nem %s:%s\n", h.hostname, h.port)
	http.ListenAndServe(addr, mux)
}