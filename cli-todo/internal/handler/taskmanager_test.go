package handler

import (
	"cli-todo/internal/model"
	"cli-todo/internal/storage"
	"errors"
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	defer os.Remove("tasks.json")
	taskList := TaskList{Tasks: model.List{}}
	taskDescription := "Task to test"

	err := taskList.AddTask(taskDescription)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	if len(taskList.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(taskList.Tasks))
	}

	var desc = taskList.Tasks[0].Description
	if desc != taskDescription {
		t.Errorf("Expected task description: \n %q \n got %q", taskDescription, desc)
	}

	os.Remove("tasks.json")
}

func TestListTasks(t *testing.T) {
	defer os.Remove("tasks.json")
	taskList := TaskList{}
	err := taskList.AddTask("First task")
	if err != nil {
		t.Fatalf("Setup failed: returned error %v", err)
	}

	tasks, err := ListTasks()
	if err != nil {
		t.Fatalf("ListTasks failed: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected at least 1 task, got %d", len(tasks))
	}

	os.Remove("tasks.json")
}

func TestCompleteTask(t *testing.T) error {
	defer os.Remove("tasks.json")

	tasklist := TaskList{Tasks: model.List{}}
	taskDescription := "testing task"

	tasklist.AddTask(taskDescription)

	lastIndex := len(tasklist.Tasks) - 1

	newTaskId := tasklist.Tasks[lastIndex].ID
	tasklist.CompleteTask(newTaskId)

	// reload
	tasks, err := storage.LoadTasks()
	if err != nil {
		t.Errorf("Can't load tasks: %v", err)
	}

	found := false
	// _ is used when not necessary in loop
	for _, task := range tasks {
		if task.ID == newTaskId {
			found = true
			if !task.Completed {
				t.Errorf("Task %v was not marked as completed", task.ID)
			}
			break
		}
	}
	if !found {
		return errors.New("Task not found")
	}

	return nil
}
