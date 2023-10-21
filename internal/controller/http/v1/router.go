package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"refactoring/internal/usecase"
	"time"
)

func NewRouter(usersUC *usecase.UsersUseCase) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Healthcheck
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	// API
	router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			v1Routes(router, usersUC)
		})
	})

	return router
}

func v1Routes(router chi.Router, usersUC *usecase.UsersUseCase) {
	setupUsersRoutes(router, usersUC)
}
