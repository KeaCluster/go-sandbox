package model

import (
	"errors"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   time.Time
}

type List []Task

func (t *Task) validate() error {
	if t.Description == "" {
		return errors.New("Description can't be empty")
	}
	return nil
}
