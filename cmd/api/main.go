package main

import (
	"github.com/lauragrassig/golang-api/internal/config"
	"github.com/lauragrassig/golang-api/internal/handler"
	"github.com/lauragrassig/golang-api/internal/httpserver"
	"github.com/lauragrassig/golang-api/internal/infra/database"
	"github.com/lauragrassig/golang-api/internal/repository"
	"github.com/lauragrassig/golang-api/internal/service"
)

func main() {
	// Config init
	env := config.Init()

	// Database: DB
	db := database.New(env.PostgresHostname, env.PostgresPort, env.PostgresUser, env.PostgresPassword, env.PostgresDatabase)

	// Repository: DB Access
	ur := repository.NewUserRepository(db.Db)

	// Service: Business Rules
	us := service.NewUserService(ur)

	// Handler: Routes
	hdl := handler.New(us);

	srv := httpserver.New(hdl, "", env.ApiPort)
	srv.Start()
}