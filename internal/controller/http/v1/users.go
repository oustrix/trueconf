package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/internal/entity"
	"refactoring/internal/usecase"
)

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
	users := u.uc.GetUsers()

	render.JSON(w, r, *users)
}

func (u *usersRoutes) createUser(w http.ResponseWriter, r *http.Request) {
	request := entity.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id := u.uc.CreateUser(request)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (u *usersRoutes) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := u.uc.GetUser(id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	} else {
		render.JSON(w, r, user)
	}
}

func (u *usersRoutes) updateUser(w http.ResponseWriter, r *http.Request) {
	request := entity.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if err := u.uc.UpdateUser(id, request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	} else {
		render.Status(r, http.StatusNoContent)
	}
}

func (u *usersRoutes) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := u.uc.DeleteUser(id); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	} else {
		render.Status(r, http.StatusNoContent)
	}

}
