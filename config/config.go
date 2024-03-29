package config

import "github.com/spf13/viper"

type Config struct {
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_DATABASE string
}

var ENV *Config

func LoadCOnfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
