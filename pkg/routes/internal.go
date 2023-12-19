package routes

import (
	"backend-manager/pkg/controllers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterInternalRoutes(router *echo.Group) {
	fmt.Print("this is route")

	router.GET("/test", controllers.HelloWorldController).Name = "testing"
	//router.GET("/pending/count/:wallet_number/", controllers.RequestMoneyPendingCountController).Name = "request-money-pending-count-internal"
}
