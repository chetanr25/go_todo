package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/chetanr25/go_todo/models"
	"github.com/chetanr25/go_todo/utils"
	"github.com/spf13/cobra"
)

func AddToList(cmd *cobra.Command, args []string) {
	newID := utils.GetNextID()
	newTitle := strings.Join(args, " ")
	newItem := models.TodoItem{
		ID:        newID,
		Title:     newTitle,
		Done:      false,
		DateAdded: time.Now().Format("2006-01-02 Monday 15:04:05"),
	}

	err := utils.WriteToCSV(newItem)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	fmt.Println("Todo item added successfully!")
}

var AddCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo item",
	Args:  cobra.MinimumNArgs(1),
	Run:   AddToList,
}
