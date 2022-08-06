package main

import (
	"github.com/Salam4nder/todo/cmd"
	"github.com/Salam4nder/todo/task"
)

func main() {
	cmd.Execute()
}

func init() {
	task.TodoList = &task.List{}
	// Since os.WriteFile truncates / overwrites the file it's writing to, it's important to
	// load the List with data on init to be able to save multiple todos to the file.
	err := task.TodoList.Load(task.FileName)
	if err != nil {
		// If the app is run for the first time, a json file to store the todos is created.
		_ = task.TodoList.Add("I am a sample todo. Delete me with the rm command: todo rm 0")
		err = task.TodoList.Write(task.FileName)
		if err != nil {
			panic(err)
		}
	}
}
