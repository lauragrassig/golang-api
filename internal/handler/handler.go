package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/lauragrassig/golang-api/common"
	"github.com/lauragrassig/golang-api/internal/schema"
	"github.com/lauragrassig/golang-api/internal/service"
)

type Handler struct {
	us *service.UserService
}

type homeResponse struct {
	Status int `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func New(us *service.UserService) *Handler {
	return &Handler{
		us: us,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	res := homeResponse{
		Status: 200,
		ErrorMessage: "This the Home",
	}

	common.WriteJson(w, res)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		common.WriteStatus(w, http.StatusBadRequest)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		common.WriteStatus(w, http.StatusBadRequest)
	}
	
	user, err := h.us.GetUser(intId)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
	}

	common.WriteJson(w, user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.us.GetUsers()

	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	common.WriteJson(w, users)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		log.Println("empty body")
		common.WriteStatus(w, http.StatusBadRequest)
		return 
	}

	var cur schema.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&cur)

	if err != nil {
		log.Println("json")
		common.WriteStatus(w, http.StatusBadRequest)
		return
	}

	err = h.us.CreateUser(&cur)

	if err != nil {
		log.Printf("error creating user: %v", err)
		common.WriteStatus(w, http.StatusInternalServerError)
		return
	}


	common.WriteStatus(w, http.StatusOK)
}