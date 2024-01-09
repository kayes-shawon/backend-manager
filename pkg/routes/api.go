package routes

import (
	"backend-manager/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterInternalRoutes(router *echo.Group) { // Pass db as a dependency
	router.GET("/test/", controllers.HelloWorldController).Name = "testing"

	// Services
	router.POST("/services/", controllers.CreateService).Name = "create_service"
	router.GET("/services/:name/", controllers.GetServiceByName).Name = "get_service_by_name"
	router.GET("/services/", controllers.GetServices).Name = "get_services"
	router.PATCH("/services/:name/", controllers.UpdateService).Name = "update_service"
	router.DELETE("/services/:name/", controllers.DeleteService).Name = "delete_service"

}
