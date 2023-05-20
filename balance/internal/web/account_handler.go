package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	getaccount "github.com/celsopires1999/fc-ms-balance/internal/usecase/get_account"
	"github.com/go-chi/chi/v5"
)

type WebAccountHandler struct {
	GetAccountUseCase getaccount.GetAccountUseCase
}

func NewWebAccountHandler(getAccountUseCase getaccount.GetAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		GetAccountUseCase: getAccountUseCase,
	}
}

func (h *WebAccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	var dto getaccount.GetAccountInputDTO
	dto.ID = chi.URLParam(r, "id")

	output, err := h.GetAccountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
