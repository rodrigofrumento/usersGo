package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodrigofrumento/usersGo/config/env"
	"github.com/rodrigofrumento/usersGo/config/logger"
	"github.com/rodrigofrumento/usersGo/internal/database"
	"github.com/rodrigofrumento/usersGo/internal/database/sqlc"
	"github.com/rodrigofrumento/usersGo/internal/handler/routes"
	"github.com/rodrigofrumento/usersGo/internal/handler/userhandler"
	"github.com/rodrigofrumento/usersGo/internal/repository/userrepository"
	"github.com/rodrigofrumento/usersGo/internal/service/userservice"
)

func main() {
	logger.InitLogger()
	slog.Info("starting API")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load env variables", err, slog.String("package", "main"))
		return
	}
	dbConnection, err := database.NewDBConn()
	if err != nil {
		slog.Error("error to connected to DB", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	//user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	//routes
	router := chi.NewRouter()
	routes.InitUserRoutes(router, newUserHandler())
	routes.InitDocsRouter(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}

}
