package models

// Session holds details of a particular session
type Session struct {
	ID     uint64
	Name   string
	Users  []User
	Status string
}
