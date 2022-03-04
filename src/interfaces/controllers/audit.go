package controllers

import (
	"net/http"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type AuditController struct {
	// -> gateway.NewAuditRepository
	RepoFactory func(c *gorm.DB) port.AuditRepository
	// -> contracts.NewAuditContracts
	ContractFactory func() port.AuditContract
	// -> presenter.NewAuditOutputPort
	OutputFactory func(w http.ResponseWriter) port.AuditOutputPort
	// -> crypt.NewAuditCrypt
	CryptFactory func(p contracts.Param) port.AuditCrypt
	// -> interactor.NewAuditInputPort
	InputFactory func(
		o port.AuditOutputPort,
		u port.AuditRepository,
		co port.AuditContract,
		cr port.AuditCrypt,
	) port.AuditInputPort
	Conn  *gorm.DB
	Param contracts.Param
}

func LoadAuditController(db *gorm.DB, param contracts.Param) *AuditController {
	return &AuditController{Conn: db, Param: param}
}

func (cc *AuditController) Post(w http.ResponseWriter, r *http.Request) {
	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	contract := cc.ContractFactory()
	crypt := cc.CryptFactory(cc.Param)
	inputPort := cc.InputFactory(outputPort, repository, contract, crypt)
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