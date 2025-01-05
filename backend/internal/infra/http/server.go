package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer(app *app.App) *http.Server {
	router := chi.NewRouter()

	// Add middlewares
	router.Use(loggingMiddleware)
	router.Use(authMiddleware)

	// Register routes (delegates to a routes package or directly calls handlers)
	registerRoutes(router, app)

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

func registerRoutes(router *chi.Mux, app *app.App) {
	router.Route(
		"/api/v1", func(r chi.Router) {
			r.Mount("/users", NewUserHandler(app.UserService).Routes())
		},
	)
}
