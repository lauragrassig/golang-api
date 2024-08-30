package main

import "github.com/lauragrassig/golang-api/httpserver"

func main() {
	srv := httpserver.New("", 8080)
	srv.Start()
}