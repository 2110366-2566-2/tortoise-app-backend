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
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
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
			fmt.Println()
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}

			done := make(chan struct{})
			go func() {
				log.Println("Timeout of 5 seconds.")
				<-ctx.Done()
				done <- struct{}{}
			}()
			<-done

			log.Println("Server exiting ...")
		}()

		// Handle cleanup for the database
		cleanup()

		log.Println("Server closed.")

	}, nil
}

func buildServer(env configs.EnvVars) (*http.Server, func(), error) {
	// init the database
	db, cancel, err := database.ConnectMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}
	dbHandler := database.NewHandler(db)

	// connect to firebase
	stg, err := storage.ConnectFirebase(context.Background(), env.FIREBASE_CONFIG)
	if err != nil {
		return nil, nil, err
	}
	stgHandler := storage.NewHandler(stg)

	// set gin mode
	if env.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if env.GIN_MODE == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// init the server
	r := gin.Default()

	// set up CORS
	r.Use(CORSMiddleware(env))

	// setup the routes
	rest.SetupRoutes(r)
	apiV1.SetupRoutes(r, dbHandler, stgHandler, env)

	// create a new server
	srv := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: r,
	}

	// print the ascii art
	printASCIIArt()

	return srv, func() {
		log.Println("Closing the database ...")
		err := database.CloseMongo(db, cancel)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		log.Println("Database closed.")
	}, nil
}

func CORSMiddleware(env configs.EnvVars) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func printASCIIArt() {

	// fmt.Println("\x1b[1;31m" + `
	// ███╗░░░███╗░█████╗░███╗░░██╗  ░░██╗██╗░░░░░░██████╗░  ██╗░░░░░██╗██╗░░░██╗
	// ████╗░████║██╔══██╗████╗░██║  ░██╔╝██║░░░░░░╚════██╗  ██║░░░░░██║██║░░░██║
	// ██╔████╔██║███████║██╔██╗██║  ██╔╝░██║█████╗░█████╔╝  ██║░░░░░██║╚██╗░██╔╝
	// ██║╚██╔╝██║██╔══██║██║╚████║  ███████║╚════╝░╚═══██╗  ██║░░░░░██║░╚████╔╝░
	// ██║░╚═╝░██║██║░░██║██║░╚███║  ╚════██║░░░░░░██████╔╝  ███████╗██║░░╚██╔╝░░
	// ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝  ░░░░░╚═╝░░░░░░╚═════╝░  ╚══════╝╚═╝░░░╚═╝░░░
	// ` + "\x1b[0m")

	fmt.Println("\x1b[38;5;172m" + `
	███████╗░███████╗████████╗██████╗░█████╗░██╗░░░░░
	██╔══██╗██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██║░░░░░
	███████╔╝█████╗░░░░██║░░░██████╔╝███████║██║░░░░░
	██╔═══╝░██╔══╝░░░░░██║░░░██╔═══╝░██╔══██║██║░░░░░
	██║░░░░░███████╗░░░██║░░░██║░░░░░██║░░██║███████╗
	╚═╝░░░░░╚══════╝░░░╚═╝░░░╚═╝░░░░░╚═╝░░╚═╝╚══════╝
    ` + "\x1b[0m")

	fmt.Println("\x1b[38;5;78m" + `
	███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
	██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
	███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
	╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
	███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
	╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝																														
	` + "\x1b[0m")
}
