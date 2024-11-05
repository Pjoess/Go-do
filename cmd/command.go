package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	Add      string
	Delete   int
	Complete int
	List     bool
}

func newFlags(todos *Todos) {
	fl := Flags{}

	flag.StringVar(&fl.Add, "add", "", "Add a new todo, specify title")
	flag.IntVar(&fl.Delete, "delete", -1, "Remove a todo, specify index")
	flag.IntVar(&fl.Complete, "complete", -1, "Complete a todo, specify index")
	flag.BoolVar(&fl.List, "list", false, "List all todo's")

	flag.Parse()
	fl.Execute(todos)
}

func (fl *Flags) Execute(todos *Todos) {
	todo := *todos

	switch {
	case fl.List:
		displayTodos(todo)
	case fl.Add != "":
		todo.add(fl.Add)
	case fl.Delete >= 0:
		todo.delete(fl.Delete)
	case fl.Complete >= 0:
		todo.finish(fl.Complete)
	default:
		fmt.Println("invalid flag used, try using -help")
	}
}
