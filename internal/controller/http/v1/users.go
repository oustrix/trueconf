package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"refactoring/internal/usecase"
)

// TODO: users usecase
type usersRoutes struct {
	uc *usecase.UsersUseCase
}

func setupUsersRoutes(router chi.Router, usersUC *usecase.UsersUseCase) {
	u := &usersRoutes{
		uc: usersUC,
	}

	router.Route("/users", func(router chi.Router) {
		router.Get("/", u.getUsers)
		router.Post("/", u.createUser)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", u.getUser)
			router.Patch("/", u.updateUser)
			router.Delete("/", u.deleteUser)
		})
	})
}

func (u *usersRoutes) getUsers(w http.ResponseWriter, r *http.Request) {
	u.uc.GetUsers(w, r)

}

func (u *usersRoutes) createUser(w http.ResponseWriter, r *http.Request) {
	u.uc.CreateUser(w, r)
}

func (u *usersRoutes) getUser(w http.ResponseWriter, r *http.Request) {
	u.uc.GetUser(w, r)
}

func (u *usersRoutes) updateUser(w http.ResponseWriter, r *http.Request) {
	u.uc.UpdateUser(w, r)
}

func (u *usersRoutes) deleteUser(w http.ResponseWriter, r *http.Request) {
	u.uc.DeleteUser(w, r)
}
