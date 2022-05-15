package config

import (
	"github.com/spf13/viper"
)

type cfgSetting struct {
	DBUsername     string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBName         string `mapstructure:"DB_NAME"`
	ATSecret       string `mapstructure:"AT_SECRET"`
	RTSecret       string `mapstructure:"RT_SECRET"`
	ETSecret       string `mapstructure:"ET_SECRET"`
	MgDomain       string `mapstructure:"MG_DOMAIN"`
	MgPublicAPIKey string `mapstructure:"MG_PUBLIC_API_KEY"`
	MgAPIKey       string `mapstructure:"MG_API_KEY"`
	MgEmailTo      string `mapstructure:"MG_EMAIL_TO"`
}

var Cfg *cfgSetting

func LoadConfig(path string) {
	// Initialise viper with config path
	viper.AddConfigPath(path)

	// Tell viper what config file name to look out for
	viper.SetConfigName("config")

	// Make sure viper overrides the environment variables with
	viper.AutomaticEnv()

	// Read the config values
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var configObject cfgSetting
	err = viper.Unmarshal(&configObject)
	if err != nil {
		panic(err)
	}

	Cfg = &configObject
}

func Init() {
	LoadConfig("./config")
}
