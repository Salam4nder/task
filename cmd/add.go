package cmd

import (
	"github.com/Salam4nder/todo/task"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the todo list.",
	Long:  "Add a task to the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		err := task.TodoList.Add(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}
		err = task.TodoList.Write(task.FileName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
