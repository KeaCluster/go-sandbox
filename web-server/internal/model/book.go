package model

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
}
