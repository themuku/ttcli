package commands

import (
	"encoding/json"
	"os"
	"taskTrackerCLI/utils"
)

func DeleteTask(order int) {
	tasks := utils.GetTasksFromJson("tasks.json")

	idx, _ := utils.GetTaskByOrder(order)

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	updatedJson, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", updatedJson, 0644)
	if err != nil {
		panic(err)
	}
}
