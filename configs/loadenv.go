package configs

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
		log.Info("ENV loaded from AutomaticEnv()")
	} else {
		log.Info("ENV loaded from .env")
	}
}
