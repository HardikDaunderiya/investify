package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource string `mapstructure:"DB_SOURCE"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	PORT     string `mapstructure:"PORT"`
}

var EnvVars Config

func LoadConfig(path string) (config Config, err error) {
	// Load the config here
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read config file:", err)
	}

	viper.AutomaticEnv()

	EnvVars.PORT = viper.GetString("PORT")
	EnvVars.DBSource = viper.GetString("DB_SOURCE")
	EnvVars.DBDriver = viper.GetString("DB_DRIVER")

	return
}
