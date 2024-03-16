package handler

import (
	"cli-todo/internal/model"
	"cli-todo/internal/storage"
	"errors"
	"fmt"
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
		return fmt.Errorf("Failed to sake task: %v", err)
	}

	tl.Tasks = append(tl.Tasks, newTask)
	if err := storage.SaveTasks(tl.Tasks); err != nil {
		return fmt.Errorf("Failed to save task: %v", err)
	}
	return nil
}

func ListTasks() ([]model.Task, error) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tl *TaskList) CompleteTask(id uuid.UUID) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return errors.New("Failed to load tasks: %v" + err.Error())
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return errors.New("Task not found")
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		return errors.New("Failed to load tasks: %v" + err.Error())
	}
	return nil
}

func (tl *TaskList) DeleteTask(id uuid.UUID) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return errors.New("Failed to load tasks: %v" + err.Error())
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return errors.New("Task not found")
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		return errors.New("Failed to load tasks: %v" + err.Error())
	}
	return nil
}
