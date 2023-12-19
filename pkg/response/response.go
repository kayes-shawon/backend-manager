package response

import (
	"github.com/labstack/echo/v4"
)

type ReturnResponseFormat struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Lang    string      `json:"lang" default:"en"`
	Data    interface{} `json:"data,omitempty"`
}

type StructResponse struct {
	HttpStatus int
	StateCode  string
	MessageEn  string `default:""`
	MessageBn  string `default:""`
}

func (r StructResponse) ResponseFormat(data interface{}, lang string) ReturnResponseFormat {
	var message string
	if lang == "bn" {
		message = r.MessageBn
	} else {
		message = r.MessageEn
	}

	return ReturnResponseFormat{
		Code:    r.StateCode,
		Message: message,
		Lang:    lang,
		Data:    data,
	}
}

func (r StructResponse) WriteToResponse(c echo.Context, data interface{}, lang string) error {
	return c.JSON(r.HttpStatus, r.ResponseFormat(data, lang))
}
