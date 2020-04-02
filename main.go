package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spankie/pointer/models"
	"github.com/spankie/pointer/router"
)

func main() {
	handler := router.NewRouter()
	handler.SetContentType("application/json")
	handler.POST("/session", func(w http.ResponseWriter, r *http.Request) {
		var session models.Session
		// get the session and save to the database
		err := json.NewDecoder(r.Body).Decode(&session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session.ID = 13
		fmt.Println(session)
		js, err := json.Marshal(session)
		if err != nil {
			fmt.Printf("Error while marshaling json: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{error: 'internal server error'}"))
			return
		}
		w.Write(js)
	})
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	gracefulShutdown(server)
	fmt.Printf("server started at %s\n", server.Addr)
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
