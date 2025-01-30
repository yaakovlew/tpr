package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string
	Port        string
	Destination string
	Test        string
	Lab2AppUrl  string
	DB          DBConfig
}

type DBConfig struct {
	Host     string
	SSLMode  string
	Port     string
	Username string
	DBName   string
	Password string
}

var AppConfig Config

func InitConfig() error {
	// Load .env file (optional, will use system env if not found)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load config from env variables
	AppConfig = Config{
		Host:        getEnv("APP_HOST", "localhost"),
		Port:        getEnv("APP_PORT", "8000"),
		Destination: getEnv("DESTINATION", "./materials/"),
		Test:        getEnv("TEST_PATH", "./src/test"),
		Lab2AppUrl:  getEnv("LAB2_APP_URL", "http://localhost:8002/lab2"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "postgres"),
			DBName:   getEnv("DB_NAME", "mephisrw"),
			Password: getEnv("DB_PASSWORD", ""), // Now password is inside env
		},
	}

	return nil
}

// Helper function to get env variable or fallback to default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
