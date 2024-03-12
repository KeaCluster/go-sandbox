package cmd

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addTask := flag.String("add", "", "Add a new task")
	listTask := flag.Bool("list", false, "List all tasks")
	completeTask := flag.Int("complete", 0, "Mark task as complete")
	deleteTask := flag.Int("delete", 0, "Delete a task")

	flag.Parse()

	switch {
	case *addTask != "":
		fmt.Println("Adding task:", *addTask)
	case *listTask:
		fmt.Println("Listing all tasks", *listTask)
	case *completeTask > 0:
		fmt.Println("Task marked as completed", *completeTask)
	case *deleteTask > 0:
		fmt.Println("Task deleted", *deleteTask)
	default:
		fmt.Println("No valid Task specified")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
