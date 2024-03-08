package cmd

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addAction := flag.String("add", "", "Add a new task")
	listAction := flag.Bool("list", false, "List all tasks")
	completeAction := flag.Int("complete", 0, "Mark task as complete")
	deleteAction := flag.Int("delete", 0, "Delete a task")

	flag.Parse()

	switch {
	case *addAction != "":
		fmt.Println("Adding task:", *addAction)
	case *listAction:
		fmt.Println("Listing all tasks", *listAction)
	case *completeAction > 0:
		fmt.Println("Task marked as completed", *completeAction)
	case *deleteAction > 0:
		fmt.Println("Task deleted", *deleteAction)
	default:
		fmt.Println("No valid action specified")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
