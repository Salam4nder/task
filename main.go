package main

import (
	"github.com/Salam4nder/todo/cmd"
	"github.com/Salam4nder/todo/task"
	"log"
)

func main() {
	cmd.Execute()
}

func init() {
	cmd.TodoList = &task.List{}
	err := cmd.TodoList.Load(".todo.json")
	if err != nil {
		log.Fatal(err)
	}
}
