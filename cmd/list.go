package commands

import (
	"github.com/chetanr25/go_todo/models"
	"github.com/chetanr25/go_todo/utils"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var todoList = []models.TodoItem{}

func DisplayList(cmd *cobra.Command, args []string) {
	utils.LoadDataFromCSV(&todoList)
	tbl := table.New("ID", "Title", "Done", "DateAdded")
	for _, item := range todoList {
		tbl.AddRow(item.ID, item.Title, item.Done, item.DateAdded)
	}
	tbl.Print()
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Run:   DisplayList,
}
