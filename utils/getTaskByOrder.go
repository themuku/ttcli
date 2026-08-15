package utils

import (
	"fmt"
	"os"
	"taskTrackerCLI/types"
)

func GetTaskByOrder(order int) (int, types.Task) {
	tasks := GetTasksFromJson("tasks.json")
	var foundTask types.Task
	var foundTaskIdx int

	for idx, task := range tasks {
		if task.Order == order {
			foundTask = task
			foundTaskIdx = idx
		}
	}

	if foundTask.Id == "" {
		err := fmt.Errorf("task with order `%s` was not found", os.Args[2])
		panic(err)
	}

	return foundTaskIdx, foundTask
}
