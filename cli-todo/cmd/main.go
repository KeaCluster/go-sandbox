package cmd

import (
	"cli-todo/internal/handler"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-todo",
	Short: "CLI Todo is a simple task manager (test)",
}

func main() {
	taskList := &handler.TaskList{}

	var addCmd = &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			description := args[0]
			err := taskList.AddTask(description)
			if err != nil {
				fmt.Printf("Error adding task %s", err)
				os.Exit(1)
			}
			fmt.Println("Task added successfully")
		},
	}

	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
