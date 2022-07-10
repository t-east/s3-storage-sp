package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sp/src/domains/entities"
	"sp/src/usecases/interactor"

	"github.com/labstack/echo/v4"
)

type AuditHandler struct {
	AuditUC *interactor.AuditUseCase
}

func NewAuditHandler(auditUC *interactor.AuditUseCase) *AuditHandler {
	return &AuditHandler{AuditUC: auditUC}
}

func (ah *AuditHandler) Post(c echo.Context) error {
	// proofs, err := ah.AuditUC.ProofGen()
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	raw, err := ioutil.ReadFile("./proof.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc entities.Proof
	json.Unmarshal(raw, &fc)
	return c.JSON(http.StatusOK, fc)
}
