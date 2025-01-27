package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host        string   `mapstructure:"host"`
	Port        string   `mapstructure:"port"`
	Destination string   `mapstructure:"destination"`
	Test        string   `mapstructure:"test"`
	Lab2AppUrl  string   `mapstructure:"lab2_app_url"`
	DB          DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	SSLMode  string `mapstructure:"sslmode"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	DBName   string `mapstructure:"dbname"`
}

var AppConfig Config

func InitConfig() error {
	// Set the path and file name for the configuration file
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // Specify that the config file is in YAML format

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil
}
