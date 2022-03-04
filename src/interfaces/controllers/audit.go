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
	// -> crypt.NewAuditCrypt
	CryptFactory func() port.AuditContract
	// -> presenter.NewAuditOutputPort
	OutputFactory func(w http.ResponseWriter) port.AuditOutputPort
	// -> interactor.NewAuditInputPort
	InputFactory func(
		o port.AuditOutputPort,
		u port.AuditRepository,
		co port.AuditContract,
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
	inputPort := cc.InputFactory(outputPort, repository, contract)
	inputPort.Challen()
}
