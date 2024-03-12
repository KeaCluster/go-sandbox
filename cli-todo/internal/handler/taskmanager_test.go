package handler

import (
	"cli-todo/internal/model"
	"os"
	"testing"
)

func TextAddTask(t *testing.T) {
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
