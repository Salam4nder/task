/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Salam4nder/todo/task"
	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Lists all todos and their dates",
	Long:  "Lists all todos and their dates",
	Run: func(cmd *cobra.Command, args []string) {
		task.TodoList.PrintWithDate()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
