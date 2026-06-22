package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrajaktaKambale/students-api/internal/config"
	"github.com/PrajaktaKambale/students-api/internal/http/handlers/student"
)

func main() {

	//load config
	cfg := config.Load()

	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", student.New())

	//setup server
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("server  started on %s", slog.String("address", cfg.Address))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to start server:", slog.String("error", err.Error()))
		}
	}()
	<-done

	slog.Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Server shutdown error:", slog.String("error", err.Error()))
	}

	fmt.Println("Server exited")

}
