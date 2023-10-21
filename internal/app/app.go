package app

import (
	"net/http"
	v1 "refactoring/internal/controller/http/v1"
	usersRepository "refactoring/internal/repository/json"
	"refactoring/internal/usecase"
)

const store = `users.json`

func Run() {
	usersRepo := usersRepository.NewUsersRepository(store)

	usersUC := usecase.NewUsersUseCase(usersRepo)

	router := v1.NewRouter(usersUC)

	http.ListenAndServe(":3333", router)
}
