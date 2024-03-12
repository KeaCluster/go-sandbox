package handler

import (
	"cli-todo/internal/model"
	"cli-todo/internal/storage"
	"errors"
	"time"

	"github.com/google/uuid"
)

type TaskList struct {
	Tasks model.List
}

func (tl *TaskList) AddTask(description string) error {
	newTask := model.Task{
		ID:          uuid.New(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	if err := newTask.Validate(); err != nil {
		return errors.New("Failed to sake task: %v" + err.Error())
	}

	tl.Tasks = append(tl.Tasks, newTask)

	if err := storage.SaveTasks(tl); err != nil {
		return errors.New("Failed to save task: %v" + err.Error())
	}
	return nil
}

func ListTasks() ([]Task, error) {
}

func CompleteTask(id int) error {
}

func DeleteTask(id int) error {
}
