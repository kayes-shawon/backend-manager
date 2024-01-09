package configs

import (
	"backend-manager/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func EchoInitialize(e *echo.Echo) {

	// request id middleware
	e.Use(middleware.RequestID())

	// adding trailing slash middleware
	e.Pre(middleware.AddTrailingSlash())

	// recover middleware
	DefaultRecoverConfig := middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger := GetRequestLogger(c)
			logger.Error().Interface("err", err).Msg("Recovered from fatal panic error.")
			return response.TechnicalError.WriteToResponse(c, struct{}{}, "en")
		},
	}
	e.Use(middleware.RecoverWithConfig(DefaultRecoverConfig))

	// request response logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"x-request-id":"${id}", "time":"${time_rfc3339_nano}", "method":"${method}", "uri":"${uri}", ` +
			`"status":${status}, "latency":"${latency_human}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	// validator config
	v := &CustomValidator{Validator: validator.New()}
	e.Validator = v

	// echo config
	e.HideBanner = viper.GetBool("HIDE_BANNER")
	e.HidePort = viper.GetBool("HIDE_PORT")
}
