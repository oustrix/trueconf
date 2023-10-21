package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	usersRepository "refactoring/internal/repository/json"
	"refactoring/internal/usecase"
	"time"
)

const store = `users.json`

var usersRepo = usersRepository.NewUsersRepository(store)

var usersUC = usecase.NewUsersUseCase(usersRepo)

func Run() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", usersUC.GetUsers)
				r.Post("/", usersUC.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", usersUC.GetUser)
					r.Patch("/", usersUC.UpdateUser)
					r.Delete("/", usersUC.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
