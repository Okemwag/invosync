package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Okemwag/invosync/internal/apps/server"
)

func main() {
	app := server.NewApp()

	// Run the application in a separate goroutine
	go func() {
		app.Start()
	}()

	// Set up signal handling to gracefully shut down the application
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan // Wait for a termination signal
	log.Println("Shutting down application...")

	// Call the shutdown method to clean up resources
	app.Shutdown()

	log.Println("Application stopped.")
}
