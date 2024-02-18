package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rodrigofrumento/usersGo/docs/custom"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	docsURL = "http://localhost:8080/docs/doc.json"
)

// @title Swagger Dark Mode
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func InitDocsRouter(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
		httpSwagger.AfterScript(custom.CustomJS),
		httpSwagger.DocExpansion("none"),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelsExpandDepth": `"-1"`,
		}),
	))
}
