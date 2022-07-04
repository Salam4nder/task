package cmd

import (
	"github.com/Salam4nder/todo/task"
	"github.com/spf13/cobra"
	"log"
)

// refCmd represents the ref command
var refCmd = &cobra.Command{
	Use:   "ref",
	Short: "Refresh the todo list by removing the tasks that are done",
	Long:  "Refresh the todo list by removing the tasks that are done",
	Run: func(cmd *cobra.Command, args []string) {
		err := task.TodoList.RemoveAllCompleted()
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
	rootCmd.AddCommand(refCmd)
}
