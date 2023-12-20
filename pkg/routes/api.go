package routes

import (
	"backend-manager/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterInternalRoutes(router *echo.Group) {
	router.GET("/test/", controllers.HelloWorldController).Name = "testing"
}
