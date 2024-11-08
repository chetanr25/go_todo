package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/chetanr25/go_todo/models"
)

func LoadDataFromCSV(todoList *[]models.TodoItem) {
	file, err := os.Open("data/data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		done, _ := strconv.ParseBool(record[2])
		*todoList = append(*todoList, models.TodoItem{
			ID:        id,
			Title:     record[1],
			Done:      done,
			DateAdded: record[3],
		})
	}
}

func RewriteCSV(items []models.TodoItem) error {
	file, err := os.Create("data/data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range items {
		err := writer.Write([]string{
			strconv.Itoa(item.ID),
			item.Title,
			strconv.FormatBool(item.Done),
			item.DateAdded,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteToCSV(item models.TodoItem) error {
	file, err := os.OpenFile("data/data.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		strconv.Itoa(item.ID),
		item.Title,
		strconv.FormatBool(item.Done),
		item.DateAdded,
	})
	return err
}

func GetNextID() int {
	var todoList []models.TodoItem
	LoadDataFromCSV(&todoList)
	if len(todoList) == 0 {
		return 1
	}
	return todoList[len(todoList)-1].ID + 1
}
