package main

import (
	"log"
	"sp/src/core"
	"sp/src/drivers/router"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"sp/src/interfaces/gateways"
	"sp/src/usecases/interactor"

	"github.com/joho/godotenv"
)

func realMain() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}
	// param, err := ethereum.GetParam()
	param, _, err := core.CreateParamMock()
	if err != nil {
		log.Fatal(err)
	}
	contentContract := contracts.NewContentContracts()
	contentRepo := gateways.NewContentRepository()
	contentCrypt := crypt.NewContentCrypt(param)
	auditContract := contracts.NewAuditContracts()
	auditCrypt := crypt.NewAuditCrypt(param)
	cu := interactor.NewContentUseCase(contentContract, contentRepo, contentCrypt)
	au := interactor.NewAuditUseCase(auditContract, auditCrypt, contentRepo)

	e := router.NewServer(cu, au)
	e.Logger.Fatal(e.Start(":4001"))
}

func main() {
	realMain()
}
