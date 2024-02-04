package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest/apiV1"
	"github.com/gin-gonic/gin"
)

func Run(env config.EnvVars) (func(), error) {
	// build the server
	srv, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// return a function to close the server and database
	return func() {
		// Handle cleanup for the server
		func() {
			// Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 5 seconds.
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}

			done := make(chan struct{})
			go func() {
				log.Println("timeout of 5 seconds.")
				<-ctx.Done()
				done <- struct{}{}
			}()
			<-done

			log.Println("Server exiting ...")
		}()

		// Handle cleanup for the database
		cleanup()
	}, nil
}

func buildServer(env config.EnvVars) (*http.Server, func(), error) {
	// init the database
	db, err := database.ConnectMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}
	handler := database.NewHandler(db)

	// init the server
	r := gin.Default()

	// setup the routes
	rest.SetupRoutes(r)
	apiV1.SetupRoutes(r, handler)

	// create a new server
	srv := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: r,
	}

	return srv, func() {
		log.Println("Closing the database ...")
		err := database.CloseMongo(db)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}, nil
}
