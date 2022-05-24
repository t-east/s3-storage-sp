package controllers

import (
	"encoding/json"
	"net/http"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/interfaces/gateways"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"
	"strconv"

	"gorm.io/gorm"
)

type UserController struct {
	// -> gateway.NewUserRepository
	RepoFactory func(c *gorm.DB) port.UserRepository
	// -> interactor.NewUserInputPort
	InputFactory func(
		u port.UserRepository,
	) port.UserInputPort
	Conn *gorm.DB
}

func LoadUserController(db *gorm.DB) *UserController {
	return &UserController{Conn: db}
}

func (uc *UserController) Post(w http.ResponseWriter, r *http.Request) {
	userReq := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repository := gateways.NewUserRepository(uc.Conn)
	inputPort := interactor.NewUserInputPort(repository)
	user, err := inputPort.Create(userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	_, tail := core.ShiftPath(r.URL.Path)
	_, tail = core.ShiftPath(tail)
	idStr, _ := core.ShiftPath(tail)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	repository := gateways.NewUserRepository(uc.Conn)
	inputPort := interactor.NewUserInputPort(repository)
	user, err := inputPort.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (uc *UserController) Dispatch(w http.ResponseWriter, r *http.Request) {
	a, _ := core.ShiftPath(r.URL.Path)
	switch r.Method {
	case "POST":
		uc.Post(w, r)
	case "GET":
		if a == "" {
			return
		}
		uc.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}
