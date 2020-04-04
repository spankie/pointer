package main

import (
	"github.com/spankie/pointer/db"
	"github.com/spankie/pointer/router"
	"github.com/spankie/pointer/server"
)

func main() {
	s := &server.Server{
		DB: db.DB{
			Sessions: make(db.Sessions, 1),
			Users:    make(db.Users, 1),
		},
		Router: router.NewRouter(),
	}
	s.Start()
}
