package routes

import (
	"github.com/labstack/echo/v4"
)

var RegisterRoutes = func(e *echo.Echo) {
	//base prefix for routes
	basePrefix := e.Group("/backend-manager")

	// route groups
	apiGroup := basePrefix.Group("/api/v1")

	RegisterInternalRoutes(apiGroup)

}
