package config

import (
	"github.com/spf13/viper"
	"log"
)

func Load() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig returns error: %s", err.Error())
	}

	viper.AutomaticEnv()
}

type AppConfig struct {
	Env    string
	Port   int
	ApiKey string
}

type DatabaseConfig struct {
	DbUrl string
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Env:    viper.GetString("ENV"),
		Port:   viper.GetInt("PORT"),
		ApiKey: viper.GetString("API_KEY"),
	}
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DbUrl: viper.GetString("DB_URL"),
	}
}
