package app

import (
	"net/http"
	"refactoring/config"
	v1 "refactoring/internal/controller/http/v1"
	usersRepository "refactoring/internal/repository/json"
	"refactoring/internal/usecase"
)

func Run(cfg *config.Config) {
	usersRepo := usersRepository.NewUsersRepository(cfg.Store.Users)

	usersUC := usecase.NewUsersUseCase(usersRepo)

	router := v1.NewRouter(usersUC)

	http.ListenAndServe(":"+cfg.HTTP.Port, router)
}
