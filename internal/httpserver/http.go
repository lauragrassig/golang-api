package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lauragrassig/golang-api/internal/handler"
)

type httpSrv struct{
	handler *handler.Handler
	hostname string
	port string
}

func New(h * handler.Handler, hostname string, port uint) *httpSrv{
	strPort := strconv.Itoa(int(port))

	return &httpSrv{
		handler: h,
		hostname: hostname,
		port: strPort,
	}
}

func (h *httpSrv) Start() {
	// addr := h.hostname + ":" + h.port

	mux := http.NewServeMux()
	initRoutes(h.handler, mux)
	addr := fmt.Sprintf("%s:%s", h.hostname, h.port)
	fmt.Printf("Servidor http rodando nem %s:%s", h.hostname, h.port)
	http.ListenAndServe(addr, mux)
}