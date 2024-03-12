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
