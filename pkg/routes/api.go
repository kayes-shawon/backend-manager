package routes

import (
	"backend-manager/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterInternalRoutes(router *echo.Group) { // Pass db as a dependency
	router.GET("/test/", controllers.HelloWorldController).Name = "testing"
	router.POST("/services/", controllers.CreateService).Name = "create_service"
	router.GET("/services/", controllers.GetServices).Name = "get_services"
}
