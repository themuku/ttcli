package commands

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Help() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	helpText := `Task Tracker CLI — Command Line Reference
Usage: ttcli <command> [arguments] [flags]

COMMNAD	ARGUMENTS	DESCRIPTION
-------	---------	-----------
add	<description>	Add a new task with a description
update	<id/order> <description>	Update the description of an existing task
delete	<id/order>	Remove a task from the tracker
mark-in-progress	<id/order>	Change task status to 'In Progress'
mark-done	<id/order>	Change task status to 'Done'
list	[status] [flags]	Display tasks. Supports optional filters and flags
help		Show this help information

Available Status Filters for 'list':
  todo	Show only tasks that need to be done
  in-progress	Show only tasks currently being worked on
  done	Show only completed tasks

Available Presentation Flags for 'list':
  -s, --short	Compact table layout (Default option)
  -l, --long	Full table layout including IDs and timestamps

Usage Examples:
  ttcli add "Buy groceries"
  ttcli list done -l
  ttcli update 1 "Buy organic milk"
`

	_, err := fmt.Fprint(w, helpText)
	if err != nil {
		panic(err)
	}

	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
