package main

import (
	"backend-manager/configs"
	"backend-manager/pkg/database"
	"backend-manager/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// env load
	configs.LoadEnv()
	// logger initialized
	configs.LoggerConfig()
	// database connect
	database.Setup()
	// server setup
	configs.EchoInitialize(e)
	// register route
	routes.RegisterRoutes(e)
	// start server
	configs.StartServer(e)

}
