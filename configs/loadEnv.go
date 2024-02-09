package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	MONGODB_URI  string `mapstructure:"MONGODB_URI"`
	MONGODB_NAME string `mapstructure:"MONGODB_NAME"`
	FRONTEND_URL string `mapstructure:"FRONTEND_URL"`
	PORT         string `mapstructure:"PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	err = godotenv.Load("./configs/config.env") // .env path
	if err != nil {
		return EnvVars{}, err
	}

	config = EnvVars{
		MONGODB_URI:  os.Getenv("MONGODB_URI"),
		MONGODB_NAME: os.Getenv("MONGODB_NAME"),
		FRONTEND_URL: os.Getenv("FRONTEND_URL"),
		PORT:         os.Getenv("PORT"),
	}

	return config, nil
}
