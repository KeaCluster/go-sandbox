package handler

import (
	"cli-todo/internal/model"
	"cli-todo/storage"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (l *model.List) AddTask(description string) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newTask := model.Task{
		ID:          id,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	if err := model; err != nil {
		return err
	}

	*l = append(*l, newTask)

	if err := storage.SaveTasks(*l); err != nil {
		return errors.New("Failed to save task: " + err.Error())
	}
	return nil
}

func ListTasks() ([]Task, error) {
}

func CompleteTask(id int) error {
}

func DeleteTask(id int) error {
}
