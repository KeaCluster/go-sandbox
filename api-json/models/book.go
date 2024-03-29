package models

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Here should go functions to create new book and probably get
