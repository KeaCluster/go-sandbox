package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
}

type List []Task

func (t *Task) Validate() error {
	if t.Description == "" {
		return errors.New("Description can't be empty")
	}
	return nil
}
