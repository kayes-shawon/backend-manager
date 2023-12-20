package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var RegisterRoutes = func(e *echo.Echo) {
	//base prefix for routes
	basePrefix := e.Group("/backend-manager")

	basePrefix.GET("/health/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "I'm ok")
	})

	// route groups
	apiGroup := basePrefix.Group("/api/v1")

	RegisterInternalRoutes(apiGroup)

}
