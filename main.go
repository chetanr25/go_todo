// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/rodaine/table"
// 	"github.com/spf13/cobra"
// )

// var toDoList = []todoListObject{
// 	{
// 		ID:        1,
// 		Title:     "Learn Go",
// 		Done:      false,
// 		DateAdded: "2021-07-01",
// 	},
// }

// func displayList(cmd *cobra.Command, args []string) {
// 	dataToList()
// 	tbl := table.New("ID", "Title", "Done", "DateAdded")
// 	for _, widget := range toDoList {
// 		tbl.AddRow(widget.ID, widget.Title, widget.Done, widget.DateAdded)
// 	}

// 	tbl.Print()
// }

// func addToList(cmd *cobra.Command, args []string) {
// 	newID := len(toDoList) + 1
// 	newTitle := strings.Join(args, " ")
// 	newItem := todoListObject{
// 		ID:        newID,
// 		Title:     newTitle,
// 		Done:      false,
// 		DateAdded: time.Now().Local().Format("2006-01-02"), // You might want to set this to the current date
// 	}
// 	toDoList = append(toDoList, newItem)
// 	fmt.Println("Added new item to the list", newItem.Title)
// 	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	err = writer.Write([]string{
// 		strconv.Itoa(newItem.ID),
// 		newItem.Title,
// 		strconv.FormatBool(newItem.Done),
// 		newItem.DateAdded,
// 	})
// 	if err != nil {
// 		fmt.Println("Error writing to CSV:", err)
// 	}
// }

// func dataToList() []todoListObject {
// 	file, err := os.Open("data.csv")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil

// 	}
// 	reader := csv.NewReader(file)
// 	data, _ := reader.ReadAll()
// 	for _, item := range data {
// 		toDoList = append(toDoList, todoListObject{
// 			ID: func() int {
// 				id, err := strconv.Atoi(item[0])
// 				if err != nil {
// 					fmt.Println("Error converting ID:", err)
// 					return 0
// 				}
// 				return id
// 			}(),
// 			Title:     item[1],
// 			Done:      item[2] == "true",
// 			DateAdded: item[3],
// 		})
// 	}

// 	return toDoList
// }

// func main() {
// 	var rootCmd = &cobra.Command{Use: "todo"}
// 	table.DefaultHeaderFormatter = func(format string, toDoList ...interface{}) string {
// 		return strings.ToUpper(fmt.Sprintf(format, toDoList...))
// 	}

// 	dataToList()

//		// defer file.Close()
//		var listCmd = &cobra.Command{
//			Use:   "display",
//			Short: "List all todo items",
//			Run:   displayList,
//		}
//		var addCmd = &cobra.Command{
//			Use:   "add",
//			Short: "Add a todo item",
//			Run:   addToList,
//		}
//		rootCmd.AddCommand(listCmd)
//		rootCmd.AddCommand(addCmd)
//		rootCmd.Execute()
//		// name := flag.String("name", "world", "The name to greet.")
//		// flag.Parse()
//		// fmt.Printf("Hello, %s!\n", *name)
//	}
package main

import (
	commands "github.com/chetanr25/go_todo/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "todo"}
	rootCmd.AddCommand(commands.ListCmd)
	rootCmd.AddCommand(commands.AddCmd)
	rootCmd.AddCommand(commands.DeleteCmd)
	rootCmd.Execute()
}
