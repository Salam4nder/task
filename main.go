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
	err := cmd.TodoList.Load(".todo.json")
	if err != nil {
		if err == os.ErrNotExist {
			_ = cmd.TodoList.Add("Used to create the json file. Content won't be saved.")
		} else {
			log.Fatal(err)
		}
	}
}
