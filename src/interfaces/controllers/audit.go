package controllers

import (
	"net/http"
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
	proofs, err := ah.AuditUC.ProofGen()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, proofs)
}
