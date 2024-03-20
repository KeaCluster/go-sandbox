package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-todo",
	Short: "CLI Todo is a simple task manager (test)",
}

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long:  "Add a new task to you task list inside /tasks.json",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Added task: ", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
