package todo

import (
	"errors"
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
	// Order is important. Split the slice into two and join them together
	// except the value at index.
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
				return errors.New("index out of bounds")
			}
		}
	}
	return nil
}

// Complete a Task by filling in the CompletedAt date. Returns an error if index outside of bound.
func (tl *List) Complete(idx int) error {
	if idx < 0 || idx > len(*t)-1 {
		return errors.New("index out of bounds")
	}
	(*tl)[idx].CompletedAt = time.Now()
	return nil
}
