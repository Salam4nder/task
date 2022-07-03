package main

import (
	"github.com/Salam4nder/todo/cmd"
	"github.com/Salam4nder/todo/task"
	"log"
	"os"
)

func main() {
	cmd.Execute()
}

func init() {
	cmd.TodoList = &task.List{}
	// Since os.WriteFile truncates / overwrites the file it's writing to, it's important to
	// load the List with data on init to be able to save multiple todos to the file.
	err := cmd.TodoList.Load(".todo.json")
	if err != nil {
		// If the app is run for the first time, a json file to store the todos is created.
		if err == os.ErrNotExist {
			_ = cmd.TodoList.Add("Used to create the json file. Content won't be saved.")
		} else {
			log.Fatal(err)
		}
	}
}
