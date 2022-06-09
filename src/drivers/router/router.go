package router

import (
	"net/http"
	"sp/src/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(cc controllers.ContentController, ac controllers.AuditController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/", hello)
	api := e.Group("/api")
	api.POST("/content", cc.Post)
	api.GET("/content/all", cc.FindAll)

	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
