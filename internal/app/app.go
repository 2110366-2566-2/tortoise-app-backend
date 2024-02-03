package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func Run(env config.EnvVars) (func(), error) {
	// build the server
	srv, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// go func() {
	// 	r.Run(":" + env.PORT)
	// }()

	// start the server
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		func() { // Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 5 seconds.
			quit := make(chan os.Signal, 1)
			// kill (no param) default send syscanll.SIGTERM
			// kill -2 is syscall.SIGINT
			// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			// catching ctx.Done(). timeout of 5 seconds.
			done := make(chan struct{})
			go func() {
				<-ctx.Done()
				log.Println("timeout of 5 seconds.")
				done <- struct{}{}
			}()
			<-done
			log.Println("Server exiting...")
		}()
	}, nil
}

func buildServer(env config.EnvVars) (*http.Server, func(), error) {
	// init the storage
	db, err := database.ConnectMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}

	// create a new handler
	// handler := database.NewHandler(db)

	// init the server
	r := gin.Default()

	rest.SetupRoutes(r)

	// create a new server
	srv := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: r,
	}

	return srv, func() {
		database.CloseMongo(db)
	}, nil
}
