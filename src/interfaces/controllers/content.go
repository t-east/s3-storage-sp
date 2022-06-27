package controllers

import (
	"net/http"
	"sp/src/domains/entities"
	"sp/src/log"
	"sp/src/usecases/interactor"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ContentHandler struct {
	contentUC *interactor.ContentUseCase
}

func NewContentHandler(contentUC *interactor.ContentUseCase) *ContentHandler {
	return &ContentHandler{contentUC: contentUC}
}
func (ch *ContentHandler) Post(c echo.Context) error {
	logger, _ := log.NewLogger()
	req := &entities.ContentIn{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed.", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	content := &entities.ContentIn{
		Address:  req.Address,
		Content:  req.Content,
		MetaData: req.MetaData,
	}
	receipt, err := ch.contentUC.Upload(content)
	if err != nil {
		logger.Error("Failed.", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipt)
}

func (ch *ContentHandler) FindAll(c echo.Context) error {
	logger, _ := log.NewLogger()

	receipts, err := ch.contentUC.FindAll()
	if err != nil {
		logger.Error("Failed.", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipts)
}
