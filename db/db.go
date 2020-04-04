package db

import "github.com/spankie/pointer/models"

// Sessions holds a list of all session identified by thier IDs
type Sessions []models.Session

// Users holds a list of all users identified by thier IDs
type Users []models.User

// DB provides access to the different db
type DB struct {
	Sessions Sessions
	Users    Users
}
