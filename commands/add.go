package commands

import (
	"taskTrackerCLI/constants"
	"taskTrackerCLI/types"
	"taskTrackerCLI/utils"
	"time"
)

func AddTask(description string) {
	currentDate := time.Now().Format("2006-01-02:15:04")

	newTask := types.Task{
		Id:          utils.GenerateId(),
		CreatedAt:   currentDate,
		UpdatedAt:   currentDate,
		Description: description,
		Status:      utils.SetTodoStatus(constants.Todo),
	}

	if err := utils.AppendJson("tasks.json", newTask); err != nil {
		panic(err)
	}
}
