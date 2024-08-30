package httpserver

import (
	"fmt"
	"net/http"
	"strconv"
)

type httpServer struct{
	hostname string
	port string
}

func New(hostname string, port uint) *httpServer{
	strPort := strconv.Itoa(int(port))

	return &httpServer{
		hostname: hostname,
		port: strPort,
	}
}

func (h *httpServer) Start() {
	// addr := h.hostname + ":" + h.port
	addr := fmt.Sprintf("%s:%s", h.hostname, h.port)
	fmt.Printf("Servidor http rodando nem %s:%s", h.hostname, h.port)
	http.ListenAndServe(addr, nil)
}