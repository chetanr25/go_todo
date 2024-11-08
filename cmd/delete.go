package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chetanr25/go_todo/models"
	"github.com/chetanr25/go_todo/utils"
	"github.com/spf13/cobra"
)

func DeleteItem(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide an ID to delete")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ID must be a number")
		return
	}

	var todoList []models.TodoItem
	utils.LoadDataFromCSV(&todoList)

	if len(todoList) == 0 {
		fmt.Println("No todo items found")
		return
	}

	var itemToDelete models.TodoItem
	var found bool
	for _, item := range todoList {
		if item.ID == id {
			itemToDelete = item
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("No todo item found with ID %d\n", id)
		return
	}

	fmt.Printf("Todo item to delete:\nID: %d\nTitle: %s\nDone: %t\nDate Added: %s\n",
		itemToDelete.ID, itemToDelete.Title, itemToDelete.Done, itemToDelete.DateAdded)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Have you completed this task? (yes/no): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if !(strings.ToLower(response) == "yes" || strings.ToLower(response) == "y") {
		fmt.Println("Task not completed. Deletion cancelled.")
		return
	}

	fmt.Print("Are you sure you want to delete this item? (yes/no): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if !(strings.ToLower(confirm) == "yes" || strings.ToLower(confirm) == "y") {
		fmt.Println("Deletion cancelled")
		return
	}

	var newList []models.TodoItem
	for _, item := range todoList {
		if item.ID != id {
			newList = append(newList, item)
		}
	}

	err = utils.RewriteCSV(newList)
	if err != nil {
		fmt.Println("Error deleting item:", err)
		return
	}

	fmt.Println("Todo item deleted successfully!")
}

var DeleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a todo item by ID",
	Run:   DeleteItem,
}
