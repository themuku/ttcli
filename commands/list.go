package commands

import (
	"taskTrackerCLI/constants"
	"taskTrackerCLI/types"
	"taskTrackerCLI/utils"
)

func ListTasks(isShort bool, filter int) {
	tasks := utils.GetTasksFromJson("tasks.json")
	var filteredTasks []types.Task

	if filter == constants.Todo {
		for _, task := range tasks {
			if task.Status == utils.SetTodoStatus(constants.Todo) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else if filter == constants.Done {
		for _, task := range tasks {
			if task.Status == utils.SetTodoStatus(constants.Done) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else if filter == constants.InProgress {
		for _, task := range tasks {
			if task.Status == utils.SetTodoStatus(constants.InProgress) {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = tasks
	}

	utils.ListTasks(isShort, filteredTasks)

}
