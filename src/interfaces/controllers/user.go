package controllers

import (
	"encoding/json"
	"net/http"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type UserController struct {
	// -> gateway.NewUserRepository
	RepoFactory func(c *gorm.DB) port.UserRepository
	// -> crypt.NewUserCrypt
	CryptFactory func(p contracts.Param) port.UserCrypt
	// -> presenter.NewUserOutputPort
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> interactor.NewUserInputPort
	InputFactory func(
		o port.UserOutputPort,
		u port.UserRepository,
		cr port.UserCrypt,
	) port.UserInputPort
	Param contracts.Param
	Conn  *gorm.DB
}

func LoadUserController(db *gorm.DB, param contracts.Param) *UserController {
	return &UserController{Conn: db, Param: param}
}

func (uc *UserController) Post(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputPort := uc.OutputFactory(w)
	repository := uc.RepoFactory(uc.Conn)
	crypt := uc.CryptFactory(uc.Param)
	inputPort := uc.InputFactory(outputPort, repository, crypt)
	inputPort.Create(user)
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	//  TODO: idの取得をちゃんとしたやつにする
	id := "1"

	outputPort := uc.OutputFactory(w)
	repository := uc.RepoFactory(uc.Conn)
	crypt := uc.CryptFactory(uc.Param)
	inputPort := uc.InputFactory(outputPort, repository, crypt)
	inputPort.FindByID(id)
}
