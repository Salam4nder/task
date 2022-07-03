package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []*Task

// Add a Task to the List of Tasks. If Task is an empty string return an error.
func (tl *List) Add(todo string) error {
	if todo == "" {
		return errors.New("task can not be empty")
	}
	t := &Task{
		Task:        todo,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*tl = append(*tl, t)
	return nil
}

// Remove a Task from the List. If index parameter is out of bounds return an error.
func (tl *List) Remove(idx int) error {
	// Assignment is used for easier syntax when splitting the slice.
	l := *tl
	if idx < 0 || idx > len(l)-1 {
		return errors.New("index out of bounds")
	}
	// Order is important. Split the slice into two and join them together, but leave out the value at index.
	*tl = append(l[:idx], l[idx+1:]...)
	return nil
}

// RemoveAllCompleted checks the List for completed Tasks and removes them from the List. Returns an error if index
// is out of bounds.
func (tl *List) RemoveAllCompleted() error {
	for idx := range *tl {
		if (*tl)[idx].Done == true {
			err := tl.Remove(idx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Complete a Task by filling in the CompletedAt date and setting Done boolean to True.
// Returns an error if index outside of bounds.
func (tl *List) Complete(idx int) error {
	if idx < 0 || idx > len(*tl)-1 {
		return errors.New("index out of bounds")
	}
	(*tl)[idx].CompletedAt = time.Now()
	(*tl)[idx].Done = true
	return nil
}

// Write will marshal the List to a json format and write it to a file named fileName.
// Returns an error if fileName is empty, if the List is empty or if the marshalling fails.
func (tl *List) Write(fileName string) error {
	if fileName == "" {
		return errors.New("file name can't be empty")
	}
	data, err := json.MarshalIndent(*tl, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// Load will read the file fileName, unmarshal it from json to data and populate List with the contents.
// Returns an error if it fails to read the file or fails to unmarshal.
func (tl *List) Load(fileName string) error {
	if fileName == "" {
		return errors.New("file name can't be empty")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, tl)
	if err != nil {
		return err
	}
	return nil
}

func (tl *List) Print() {
	if len(*tl) <= 0 {
		fmt.Println("Your list of todos are empty. Try adding one with the add command.")
	}
	for idx, val := range *tl {
		fmt.Printf("%d, %v\n", idx, *val)
	}
}
