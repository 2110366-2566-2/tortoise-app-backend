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

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest/apiV1"
	"github.com/gin-gonic/gin"
)

func Run(env configs.EnvVars) (func(), error) {
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

func buildServer(env configs.EnvVars) (*http.Server, func(), error) {
	// init the database
	db, err := database.ConnectMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}
	handler := database.NewHandler(db)

	// init the server
	r := gin.Default()

	// set up CORS
	r.Use(CORSMiddleware(env))

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

func CORSMiddleware(env configs.EnvVars) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", env.FRONTEND_URL)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-access-token, x-refresh-token, x-user-id, x-user-role, x-user-email, x-user-name, x-user-phone, x-user-address, x-user-birthdate, x-user-g")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
