package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	commands2 "taskTrackerCLI/commands"
	"taskTrackerCLI/constants"
)

func main() {
	if len(os.Args) < 2 {
		commands2.Help()
		return
	}

	commands := []string{
		"add",
		"update",
		"delete",
		"mark-in-progress",
		"mark-done",
		"list",
		"help",
	}
	userCommand := os.Args[1]

	if !slices.Contains(commands, userCommand) {
		fmt.Println("Invalid command. Run 'ttcli help' for usage details.")
		os.Exit(1)
	}

	switch userCommand {
	case "help":
		commands2.Help()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing description for the task.")
			os.Exit(1)
		}
		commands2.AddTask(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing order or description.")
			os.Exit(1)
		}
		order, err := strconv.Atoi(os.Args[2])
		if err != nil {
			err = fmt.Errorf("order must be a number but got %v", os.Args[2])
			panic(err)
		}
		description := os.Args[3]

		commands2.UpdateTask(order, description, -1)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task order to delete.")
			os.Exit(1)
		}
		order, err := strconv.Atoi(os.Args[2])
		if err != nil {
			err = fmt.Errorf("order must be a number but got %v", os.Args[2])
			panic(err)
		}

		commands2.DeleteTask(order)
	case "list":
		isShort := true
		listCommands := []string{
			"todo",
			"done",
			"in-progress",
			"-l",
			"--long",
			"-s",
			"--short",
		}

		for _, arg := range os.Args[2:] {
			if !slices.Contains(listCommands, arg) {
				err := fmt.Errorf("invalid argument or flag '%s' for 'list' command", arg)
				fmt.Println("Run 'ttcli help' to see available options.")
				panic(err)
			}
		}

		if slices.Contains(os.Args, "--long") || slices.Contains(os.Args, "-l") {
			isShort = false
		} else if slices.Contains(os.Args, "--short") || slices.Contains(os.Args, "-s") {
			isShort = true
		}

		if slices.Contains(os.Args, "todo") {
			commands2.ListTasks(isShort, constants.Todo)
		} else if slices.Contains(os.Args, "done") {
			commands2.ListTasks(isShort, constants.Done)
		} else if slices.Contains(os.Args, "in-progress") {
			commands2.ListTasks(isShort, constants.InProgress)
		} else {
			commands2.ListTasks(isShort, -1)
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task order.")
			os.Exit(1)
		}
		order, err := strconv.Atoi(os.Args[2])
		if err != nil {
			err = fmt.Errorf("order must be a number but got %v", os.Args[2])
			panic(err)
		}
		commands2.UpdateTask(order, "", constants.Done)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task order.")
			os.Exit(1)
		}
		order, err := strconv.Atoi(os.Args[2])
		if err != nil {
			err = fmt.Errorf("order must be a number but got %v", os.Args[2])
			panic(err)
		}
		commands2.UpdateTask(order, "", constants.InProgress)
	}
}
