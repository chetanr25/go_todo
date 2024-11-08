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
