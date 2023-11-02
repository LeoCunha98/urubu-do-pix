package main

import (
	"database/sql"
	"net/http"

	"github.com/LeoCunha98/urubu-do-pix/internal/infra/repository"
	"github.com/LeoCunha98/urubu-do-pix/internal/infra/web"
	"github.com/LeoCunha98/urubu-do-pix/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/urubu-do-pix")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewUserRepositoryMySql(db)
	createUserUseCase := usecase.NewCreateUserUseCase(repository)
	listUsersUseCase := usecase.NewListUsersUseCase(repository)

	userHandlers := web.NewUserHandlers(createUserUseCase, listUsersUseCase)

	r := chi.NewRouter()
	r.Post("/users", userHandlers.CreateUserHandler)
	r.Get("/users", userHandlers.ListUsersHandler)

	go http.ListenAndServe(":8080", r)
}
