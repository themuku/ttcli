package utils

import (
	"encoding/json"
	"os"
	"taskTrackerCLI/types"
)

func AppendJson(fileName string, newTask types.Task) error {
	tasks := GetTasksFromJson(fileName)

	newTask.Order = len(tasks) + 1

	tasks = append(tasks, newTask)

	updatedJson, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	return os.WriteFile(fileName, updatedJson, 0644)
}
