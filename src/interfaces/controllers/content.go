package controllers

import (
	"encoding/json"
	"net/http"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func() port.ContentContract
	// -> crypt.NewContentCrypt
	CryptFactory func() port.ContentContract
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		co port.ContentContract,
	) port.ContentInputPort
	Conn *gorm.DB
}

func LoadContentController(db *gorm.DB, param contracts.Param) *ContentController {
	return &ContentController{Conn: db}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	content := &entities.Content{}
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	contract := cc.ContractFactory()
	inputPort := cc.InputFactory(outputPort, repository, contract)
	inputPort.Upload(content)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	//  TODO: idの取得をちゃんとしたやつにする
	id := "1"

	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	contract := cc.ContractFactory()
	inputPort := cc.InputFactory(outputPort, repository, contract)
	inputPort.FindByID(id)
}
