package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func setLevel(level string) {
	var l zerolog.Level
	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)
}

func LoggerConfig() {
	logLevel := viper.GetString("LOG_LEVEL")
	setLevel(logLevel)
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}

func GetRequestLogger(c echo.Context) zerolog.Logger {
	return log.Logger.With().Str("x-request-id", c.Response().Header().Get("X-Request-ID")).Logger()
}
