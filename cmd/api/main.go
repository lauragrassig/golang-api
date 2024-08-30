package main

import (
	"github.com/lauragrassig/golang-api/internal/handler"
	"github.com/lauragrassig/golang-api/internal/httpserver"
)

func main() {
	hdl := handler.New();

	srv := httpserver.New(hdl, "", 8080)
	srv.Start()
}