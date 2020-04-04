package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spankie/pointer/db"
	"github.com/spankie/pointer/models"
	"github.com/spankie/pointer/router"
)

type Server struct {
	DB     db.DB
	Router *router.Router
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
		// TODO: handle err
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) handleAddSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var session models.Session
		// get the session and save to the database
		err := s.decode(w, r, &session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session.ID = 13
		fmt.Println(session)
		s.respond(w, r, session, 201)
	}
}

func (s *Server) handleGetSessions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sessions []models.Session
		// get the session and save to the database
		sessions = s.DB.Sessions
		s.respond(w, r, sessions, 200)
	}
}

func (s *Server) routes() {
	s.Router.GET("/sessions", s.handleGetSessions())
	s.Router.POST("/sessions", s.handleAddSession())
}

func (s *Server) Start() {
	s.Router.SetContentType("application/json")
	s.routes()
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.Router,
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
