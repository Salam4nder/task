package cmd

import (
	"github.com/Salam4nder/todo/task"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all todos",
	Long:  "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		task.TodoList.Print()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
