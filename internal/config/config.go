package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")

	viper.BindEnv("db.host", "DATABASE_HOST")
	viper.BindEnv("db.user", "DATABASE_USER")
	viper.BindEnv("db.password", "DATABASE_PASSWORD")
	viper.BindEnv("db.port", "DATABASE_PORT")
	viper.BindEnv("db.name", "DATABASE_NAME")
	viper.BindEnv("host", "HOST")
	viper.BindEnv("port", "PORT")

	viper.AutomaticEnv()

	viper.SetDefault("db.host", "postgres")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "verySecretWord")
	viper.SetDefault("db.name", "postgres")
	viper.SetDefault("db.port", 5432)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; ignore error if desired")
		} else {
			return err
		}
	}

	return nil
}
