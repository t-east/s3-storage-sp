package router

import (
	"net/http"
	"sp/src/interfaces/controllers"
	"sp/src/usecases/interactor"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(cu *interactor.ContentUseCase, au *interactor.AuditUseCase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Pre(middleware.RemoveTrailingSlash())

	ch := controllers.NewContentHandler(cu)
	ah := controllers.NewAuditHandler(au)

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/", hello)
	api := e.Group("/api")
	api.POST("/content", ch.Post)
	api.GET("/content/all", ch.FindAll)
	api.POST("/proof", ah.Post)

	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
