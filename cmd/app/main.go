package main

import (
	"fmt"
	"os"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/app"
)

// @title PetPal API
// @version 1.0
// @description PetPal API is a simple API for pet marketplaces.
// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Authorization by API key (Format: "Bearer <API_KEY>")
func main() {
	// Set the exit code to 0 by default
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	// Load the environment variables
	env, err := configs.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Run the app
	cleanup, err := app.Run(env)

	// Close the server and database when the app is done
	defer cleanup()

	// Handle any errors
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
}
