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
	"github.com/rodrigofrumento/usersGo/internal/handler"
	"github.com/rodrigofrumento/usersGo/internal/handler/routes"
	"github.com/rodrigofrumento/usersGo/internal/repository/categoryrepository"
	"github.com/rodrigofrumento/usersGo/internal/repository/productrepository"
	"github.com/rodrigofrumento/usersGo/internal/repository/userrepository"
	"github.com/rodrigofrumento/usersGo/internal/service/categoryservice"
	"github.com/rodrigofrumento/usersGo/internal/service/productservice"
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
	dbConn, err := database.NewDBConn()
	if err != nil {
		slog.Error("failed to connected DB", "err", err, slog.String("package", "main"))
		return
	}

	queries := sqlc.New(dbConn)

	//user
	userRepo := userrepository.NewUserRepository(dbConn, queries)
	newUserService := userservice.NewUserService(userRepo)

	//category
	categoryRepo := categoryrepository.NewCategoryRepository(dbConn, queries)
	newCategoryService := categoryservice.NewCategoryService(categoryRepo)

	//product
	productRepo := productrepository.NewProductRepository(dbConn, queries)
	newProductService := productservice.NewProductService(productRepo)

	newHandler := handler.NewHandler(newUserService, newCategoryService, newProductService)

	//Routes
	router := chi.NewRouter()
	routes.InitRoutes(router, newHandler)
	routes.InitDocsRouter(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error started server", err, slog.String("package", "main"))
	}

}
