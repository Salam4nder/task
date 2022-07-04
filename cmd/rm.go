package cmd

import (
	"github.com/Salam4nder/todo/task"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task from the todo list",
	Long:  "Delete a task from the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		// Join the args arguments to a single string.
		arg := strings.Join(args, "")
		// Try and convert the string to a number
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Failed to read a number ")
		}
		err = task.TodoList.Remove(num)
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
	rootCmd.AddCommand(delCmd)
}
