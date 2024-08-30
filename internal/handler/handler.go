package handler

import (
	"net/http"

	"github.com/lauragrassig/golang-api/common"
)

type Handler struct {

}

type homeResponse struct {
	Status int `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	res := homeResponse{
		Status: 200,
		ErrorMessage: "This the Home",
	}

	common.WriteJson(w, res)
}

func (h *Handler) Usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Essa é a tela de usuarios"))
}