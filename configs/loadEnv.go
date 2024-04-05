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
}

func LoadConfig() (config EnvVars, err error) {
	err = godotenv.Load("./configs/config.env") // .env path
	if err != nil {
		return EnvVars{}, err
	}

	config = EnvVars{
		MONGODB_URI:     os.Getenv("MONGODB_URI"),
		MONGODB_NAME:    os.Getenv("MONGODB_NAME"),
		FRONTEND_URL:    os.Getenv("FRONTEND_URL"),
		PORT:            os.Getenv("PORT"),
		JWT_SECRET:      os.Getenv("JWT_SECRET"),
		STRIPE_KEY:      os.Getenv("STRIPE_KEY"),
		GIN_MODE:        os.Getenv("GIN_MODE"),
		FIREBASE_CONFIG: os.Getenv("FIREBASE_CONFIG"),
	}

	return config, nil
}
