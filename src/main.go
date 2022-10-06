package main

import (
	"log"
	"sp/src/domains/entities"
	"sp/src/drivers/ethereum"
	"sp/src/drivers/router"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"sp/src/interfaces/gateways"
	"sp/src/usecases/interactor"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func realMain() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}
	param := &entities.Param{
		Pairing: p.Pairing,
		G:       []byte(p.G),
		U:       []byte(p.U),
	}
	contentContract := contracts.NewContentContracts()
	contentRepo := gateways.NewContentRepository()
	auditContract := contracts.NewAuditContracts()
	auditCrypt := crypt.NewAuditCrypt(param)
	cu := interactor.NewContentUseCase(contentContract, contentRepo, auditCrypt)
	au := interactor.NewAuditUseCase(auditContract, contentContract, auditCrypt, contentRepo)

	e := router.NewServer(cu, au)
	e.Logger.Fatal(e.Start(":4001"))
}

func main() {
	realMain()
}
