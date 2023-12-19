package utils

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func GetLanguage(c echo.Context) string {
	lang := c.Request().Header.Get("Accept-Language")
	lang = strings.ToLower(lang)
	if lang == "bn" || lang == "ban" || lang == "bang" || lang == "bangla" || lang == "bengali" {
		return "bn"
	} else {
		return "en"
	}
}
