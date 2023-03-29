package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppSecret string   `mapstructure:"app_secret"`
	Database  Database `mapstructure:"database"`
}

type Database struct {
	Dsn string `mapstructure:"dsn"`
}

var C Config

func Init() {
	var configFileName string = "config.dev.json"
	if os.Getenv("GIN_MODE") == "release" {
		configFileName = "config.prod.json"
	}

	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName(configFileName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}
