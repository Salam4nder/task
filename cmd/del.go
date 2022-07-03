package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a task from the todo list.",
	Long:  "Delete a task from the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, "")
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(num)
		}
		err = TodoList.Remove(num)
		if err != nil {
			log.Fatal()
		}
		err = TodoList.Write(".todo.json")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
