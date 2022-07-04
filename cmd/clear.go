package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the whole list of todos.",
	Long:  "Clear the whole list of todos.",
	Run: func(cmd *cobra.Command, args []string) {
		err := TodoList.Clear()
		if err != nil {
			log.Fatal(err)
		}
		err = TodoList.Write(FileName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
