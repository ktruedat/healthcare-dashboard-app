package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ktruedat/healthy/config"
)

type App struct {
	Config     *config.Config
	DB         *db.Database
	HTTPServer *http.Server
}

// NewApp initializes the application
func NewApp() (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	// Initialize dependencies
	database, err := db.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	// Initialize HTTP server and handlers
	httpServer := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: handlers.NewRouter(database), // Pass dependencies to handlers
	}

	return &App{
		Config:     cfg,
		DB:         database,
		HTTPServer: httpServer,
	}, nil
}

// Run starts the application and listens for OS signals
func (a *App) Run() error {
	// Graceful shutdown support
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	// Start the HTTP server
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting HTTP server on", a.Config.Server.Address)
		if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	// Perform cleanup
	cleanupCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HTTPServer.Shutdown(cleanupCtx); err != nil {
		log.Printf("Error shutting down HTTP server: %v", err)
	}

	a.DB.Close() // Close database connections, if needed

	wg.Wait()
	log.Println("Application stopped.")
	return nil
}
