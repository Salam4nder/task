package cmd

import (
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
		err := TodoList.Add(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}
		err = TodoList.Write(".todo.json")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
