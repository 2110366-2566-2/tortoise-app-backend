package main

import (
	"fmt"
	"os"

	config "github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/app"
)

func main() {
	// Set the exit code to 0 by default
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	// Load the environment variables
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	// go routine

	// Run the app
	cleanup, err := app.Run(env)

	// Close the server and database
	defer cleanup()

	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
}
