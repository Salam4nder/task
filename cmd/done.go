package cmd

import (
	"github.com/Salam4nder/todo/task"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task complete",
	Long:  "Mark a task complete",
	Run: func(cmd *cobra.Command, args []string) {
		// Join the args arguments to a single string.
		arg := strings.Join(args, "")
		// Try and convert the string to a number
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Failed to read a number")
		}
		err = task.TodoList.Complete(num)
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
	rootCmd.AddCommand(doneCmd)
}
