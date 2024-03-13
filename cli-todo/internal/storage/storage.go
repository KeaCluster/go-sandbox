package storage

import (
	"cli-todo/internal/model"
	"encoding/json"
	"os"
)

const tasksPath = "tasks.json"

func SaveTasks(tasks model.List) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(tasksPath, data, 0644)
	return err
}

func LoadTasks() (model.List, error) {
	var tasks model.List
	data, err := os.ReadFile(tasksPath)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DoneTask(id int) {
}
