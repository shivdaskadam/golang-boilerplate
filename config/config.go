package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// AppConfig holds the configuration values from the environment variables
type AppConfig struct {
	// MySQL Configurations
	SQLUsername string `envconfig:"DB_USERNAME"`
	SQLPassword string `envconfig:"DB_PASSWORD"`
	SQLHost     string `envconfig:"DB_HOST"`
	SQLDBName   string `envconfig:"DB_NAME"`
	SQLPort     string `envconfig:"DB_PORT"`
	// MongoDB Configurations
	MongoURL    string `envconfig:"MONGO_URL"`
	MongoDBName string `envconfig:"MONGODB_DATABASE_NAME"`
}

var instance *AppConfig

// Init initializes the AppConfig instance by loading values from environment variables
func Init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var i AppConfig
	err = envconfig.Process("", &i)
	if err != nil {
		log.Fatal("Error loading configs from env", err.Error())
	}

	instance = &i
	log.Println("Configuration loaded:")
}

// Instance returns the AppConfig instance, initializing it if necessary
func Instance() *AppConfig {
	if instance == nil {
		Init()
	}

	return instance
}
