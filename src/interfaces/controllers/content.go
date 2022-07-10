package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	// content := &entities.ContentIn{
	// 	Address:  req.Address,
	// 	Content:  req.Content,
	// 	MetaData: req.MetaData,
	// }
	// receipt, err := ch.contentUC.Upload(content)
	// if err != nil {
	// logger.Error("Failed.", zap.Error(err))
	// return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	raw, err := ioutil.ReadFile("./content.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc entities.Receipt

	json.Unmarshal(raw, &fc)
	return c.JSON(http.StatusCreated, &entities.Receipt{
		ID:       fc.ID,
		Content:  req.Content,
		MetaData: req.MetaData,
		HashData: fc.HashData,
	})
}

func (ch *ContentHandler) FindAll(c echo.Context) error {
	// logger, _ := log.NewLogger()

	// receipts, err := ch.contentUC.FindAll()
	// if err != nil {
	// 	logger.Error("Failed.", zap.Error(err))
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	raw, err := ioutil.ReadFile("./log.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc entities.Log

	json.Unmarshal(raw, &fc)
	receipts := []entities.Log{
		fc,
		fc,
	}
	return c.JSON(http.StatusCreated, receipts)
}
