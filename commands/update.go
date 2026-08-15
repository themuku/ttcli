package commands

import (
	"encoding/json"
	"os"
	"taskTrackerCLI/utils"
)

func UpdateTask(order int, description string, status int) {
	//	get all tasks
	tasks := utils.GetTasksFromJson("tasks.json")

	//	get task by the order
	idx, _ := utils.GetTaskByOrder(order)

	if description != "" {
		tasks[idx].Description = description
	}

	if status != -1 {
		tasks[idx].Status = utils.SetTodoStatus(status)
	}

	updatedJson, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", updatedJson, 0644)
	if err != nil {
		panic(err)
	}
}

// todo: update task and id logic
