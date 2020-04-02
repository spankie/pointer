package models

// Session holds details of a particular session
type Session struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Users  []User `json:"users"`
	Status string `json:"status"`
}
