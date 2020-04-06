package server

import (
	"fmt"
	"net/http"

	"github.com/spankie/pointer/models"
)

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
