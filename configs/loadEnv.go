package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	MONGODB_URI     string `mapstructure:"MONGODB_URI"`
	MONGODB_NAME    string `mapstructure:"MONGODB_NAME"`
	FRONTEND_URL    string `mapstructure:"FRONTEND_URL"`
	PORT            string `mapstructure:"PORT"`
	JWT_SECRET      string `mapstructure:"JWT_SECRET"`
	STRIPE_KEY      string `mapstructure:"STRIPE_KEY"`
	GIN_MODE        string `mapstructure:"GIN_MODE"`
	FIREBASE_CONFIG string `mapstructure:"FIREBASE_CONFIG"`
	SKIP_WAIT       string `mapstructure:"SKIP_WAIT"`
}

func LoadConfig() (config EnvVars, err error) {
	err = godotenv.Load("./configs/config.env") // .env path
	if err != nil {
		return EnvVars{}, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	frontend_url := os.Getenv("FRONTEND_URL")
	if frontend_url == "" {
		frontend_url = "http://localhost:3000"
	}

	config = EnvVars{
		MONGODB_URI:     os.Getenv("MONGODB_URI"),
		MONGODB_NAME:    os.Getenv("MONGODB_NAME"),
		FRONTEND_URL:    frontend_url,
		PORT:            port,
		JWT_SECRET:      os.Getenv("JWT_SECRET"),
		STRIPE_KEY:      os.Getenv("STRIPE_KEY"),
		GIN_MODE:        os.Getenv("GIN_MODE"),
		FIREBASE_CONFIG: os.Getenv("FIREBASE_CONFIG"),
		SKIP_WAIT:       os.Getenv("SKIP_WAIT"),
	}

	return config, nil
}
