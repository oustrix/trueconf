package usecase

import (
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

func (uc *UsersUseCase) GetUsers() *entity.UserList {
	return uc.repo.GetUsers()
}

func (uc *UsersUseCase) GetUser(id string) (*entity.User, error) {
	user, err := uc.repo.GetUser(id)

	return &user, err
}

func (uc *UsersUseCase) CreateUser(request entity.CreateUserRequest) string {
	u := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := uc.repo.CreateUser(&u)

	return id
}

func (uc *UsersUseCase) UpdateUser(id string, request entity.UpdateUserRequest) error {

	user, err := uc.repo.GetUser(id)
	if err != nil {
		return err
	}

	user.DisplayName = request.DisplayName

	err = uc.repo.UpdateUser(id, &user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UsersUseCase) DeleteUser(id string) error {
	return uc.repo.DeleteUser(id)
}
