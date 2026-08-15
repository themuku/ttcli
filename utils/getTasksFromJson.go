package utils

import (
	"encoding/json"
	"os"
	"taskTrackerCLI/types"
)

func GetTasksFromJson(fileName string) []types.Task {
	var tasks []types.Task

	if _, err := os.Stat(fileName); err == nil {
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		if len(fileData) > 0 {
			if err := json.Unmarshal(fileData, &tasks); err != nil {
				panic(err)
			}
		}
	}

	return tasks
}
