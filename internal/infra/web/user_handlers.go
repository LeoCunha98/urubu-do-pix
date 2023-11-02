package web

import (
	"encoding/json"
	"net/http"

	"github.com/LeoCunha98/urubu-do-pix/internal/usecase"
)

type UserHandlers struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	ListUsersUseCase  *usecase.ListUsersUseCase
}

func NewUserHandlers(createUserUseCase *usecase.CreateUserUseCase, listUsersUseCase *usecase.ListUsersUseCase) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase: createUserUseCase,
		ListUsersUseCase:  listUsersUseCase,
	}
}

func (h *UserHandlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateUserUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *UserHandlers) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListUsersUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
