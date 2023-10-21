package usecase

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	v1 "refactoring/internal/controller/http/v1"
	"refactoring/internal/entity"
	repo "refactoring/internal/repository/json"
	"time"
)

type UsersUseCase struct {
	repo *repo.UsersRepository
}

func NewUsersUseCase(r *repo.UsersRepository) *UsersUseCase {
	return &UsersUseCase{
		repo: r,
	}
}

func (uc *UsersUseCase) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := uc.repo.GetUsers()

	render.JSON(w, r, users)
}

func (uc *UsersUseCase) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := uc.repo.GetUser(id)
	if err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func (uc *UsersUseCase) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := entity.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	u := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := uc.repo.CreateUser(&u)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (uc *UsersUseCase) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := entity.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	user, err := uc.repo.GetUser(id)
	if err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	user.DisplayName = request.DisplayName

	err = uc.repo.UpdateUser(id, &user)
	if err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (uc *UsersUseCase) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := uc.repo.DeleteUser(id)
	if err != nil {
		_ = render.Render(w, r, v1.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
