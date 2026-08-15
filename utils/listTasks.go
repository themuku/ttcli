package utils

import (
	"fmt"
	"os"
	"taskTrackerCLI/types"
	"text/tabwriter"
)

func ListTasks(isShort bool, tasks []types.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	_, err := fmt.Fprintln(w, "ID\tTASK DESCRIPTION\tORDER\tSTATUS\tCREATED AT\tUPDATED AT")
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(w, "--\t--------------\t-----\t------\t----------\t----------")
	if err != nil {
		return
	}

	for _, task := range tasks {
		var id string

		if isShort {
			id = task.Id[:4]
		} else {
			id = task.Id
		}

		_, err := fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n",
			id,
			task.Description,
			task.Order,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}
