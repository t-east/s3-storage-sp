package controllers

import (
	"encoding/json"
	"net/http"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"sp/src/interfaces/gateways"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type AuditController struct {
	// -> gateway.NewAuditRepository
	RepoFactory func() port.AuditRepository
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

func LoadAuditController(param *entities.Param) *AuditController {
	return &AuditController{Param: param}
}

func (ac *AuditController) Post(w http.ResponseWriter, r *http.Request) {
	repository := gateways.NewProofRepository()
	contentRepo := gateways.NewContentRepository()
	contract := contracts.NewAuditContracts()
	crypt := crypt.NewAuditCrypt(ac.Param)
	inputPort := interactor.NewAuditInputPort(contract, crypt, repository, contentRepo)
	proofs, err := inputPort.ProofGen()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(proofs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (ac *AuditController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ac.Post(w, r)
	default:
		http.NotFound(w, r)
	}
}
