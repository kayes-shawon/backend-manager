package controllers

import (
	"backend-manager/pkg/response"
	"backend-manager/pkg/utils"
	"github.com/labstack/echo/v4"
)

func HelloWorldController(c echo.Context) error {
	lang := utils.GetLanguage(c)
	responseData := map[string]string{"hello": "world"}

	return response.HelloWorld.WriteToResponse(c, responseData, lang)
}
