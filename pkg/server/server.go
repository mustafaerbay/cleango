package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
This implementation creates a Server type that takes an address and an HTTP handler as parameters. 
The Start() method starts the server in a new goroutine, and waits for an interrupt signal to 
gracefully shut down the server. When the signal is received, the Shutdown() method is 
called on the http.Server to shut down the server gracefully. 
If the shutdown is successful, the method returns nil. Otherwise, an error is returned.
*/
type Server struct {
	httpServer *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	fmt.Printf("Starting server at %s\n", s.httpServer.Addr)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	// Set a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
		return err
	}

	fmt.Println("Server shutdown complete.")

	return nil
}
