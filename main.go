package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spankie/pointer/router"
)

func main() {
	handler := router.NewRouter()
	handler.POST("/session", func(w http.ResponseWriter, r *http.Request) {
		// get the session and save to the database
		w.Write([]byte("Hello world"))
	})
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	gracefulShutdown(server)
	fmt.Println(server.ListenAndServe())
}

// gracefulShutdown shutsdown the server gracefully
// when ctrl c is sent from the command line
func gracefulShutdown(s *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			s.Shutdown(context.Background())
			fmt.Println("server shutdown successfully")
		}
	}()
}
