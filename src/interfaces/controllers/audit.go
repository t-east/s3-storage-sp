package controllers

import (
	"net/http"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"sp/src/interfaces/gateways"
	"sp/src/interfaces/storage"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type AuditController struct {
	// -> gateway.NewAuditRepository
	RepoFactory func(c *gorm.DB) port.AuditRepository
	// -> contracts.NewAuditContracts
	ContractFactory func() port.AuditContract
	// -> crypt.NewAuditCrypt
	CryptFactory func(p *entities.Param) port.AuditCrypt
	// -> interactor.NewAuditInputPort
	InputFactory func(
		u port.AuditRepository,
		co port.AuditContract,
		cr port.AuditCrypt,
	) port.AuditInputPort
	Conn  *gorm.DB
	Param *entities.Param
}

func LoadAuditController(db *gorm.DB, param *entities.Param) *AuditController {
	return &AuditController{Conn: db, Param: param}
}

func (ac *AuditController) Post(w http.ResponseWriter, r *http.Request) {
	repository := gateways.NewProofRepository(ac.Conn)
	contract := contracts.NewAuditContracts()
	crypt := crypt.NewAuditCrypt(ac.Param)
	storage := storage.NewContentStorage()
	inputPort := interactor.NewAuditInputPort(contract, crypt, storage, repository)
	inputPort.Challen()
}

func (ac *AuditController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ac.Post(w, r)
	default:
		http.NotFound(w, r)
	}
}
