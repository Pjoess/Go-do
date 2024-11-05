package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
	writeJSON(*todos)
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.checkIndex(index); err != nil {
		return err
	}

	// do functionality of deleting the todo
	*todos = append(t[:index], t[index+1:]...)
	writeJSON(*todos)

	return nil
}

func (todos *Todos) finish(index int) error {
	t := *todos

	if err := t.checkIndex(index); err != nil {
		return err
	}
	fmt.Println("index checked")
	isCompleted := t[index].Completed
	t[index].Completed = !isCompleted

	writeJSON(t)
	return nil
}

func (todos *Todos) checkIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalite index")
		fmt.Println(err)
		return err
	}
	return nil
}
