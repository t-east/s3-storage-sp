package controllers

import (
	"net/http"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/gateways"
	"sp/src/interfaces/storage"
	"sp/src/log"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func() port.ContentContract
	// -> crypt.NewContentCrypt
	CryptFactory func() port.ContentContract
	// -> interactor.NewContentInputPort
	InputFactory func(
		u port.ContentRepository,
		co port.ContentContract,
	) port.ContentInputPort
	Conn  *gorm.DB
	Param *entities.Param
}

func LoadContentController(param *entities.Param) *ContentController {
	return &ContentController{Param: param}
}

func (cc *ContentController) Post(c echo.Context) error {
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
	repository := gateways.NewContentRepository()
	contract := contracts.NewContentContracts()
	storage := storage.NewContentStorage()
	inputPort := interactor.NewContentInputPort(
		repository,
		contract,
		storage,
	)
	receipt, err := inputPort.Upload(content, cc.Param)
	if err != nil {
		logger.Error("Failed.", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipt)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	_, tail := core.ShiftPath(r.URL.Path)
	_, tail = core.ShiftPath(tail)
	id, _ := core.ShiftPath(tail)
	repository := gateways.NewContentRepository()
	contract := contracts.NewContentContracts()
	storage := storage.NewContentStorage()
	inputPort := interactor.NewContentInputPort(
		repository,
		contract,
		storage,
	)
	inputPort.FindByID(id)
}

func (cc *ContentController) FindAll(c echo.Context) error {
	logger, err := log.NewLogger()
	repository := gateways.NewContentRepository()
	contract := contracts.NewContentContracts()
	storage := storage.NewContentStorage()
	inputPort := interactor.NewContentInputPort(
		repository,
		contract,
		storage,
	)
	receipts, err := inputPort.FindAll()
	if err != nil {
		logger.Error("Failed.", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipts)
}
