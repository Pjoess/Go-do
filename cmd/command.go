package main

import (
	"errors"
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

	flag.StringVar(&fl.Add, "add", "none", "Add a new todo, specify title")
	flag.IntVar(&fl.Delete, "delete", -1, "Remove a todo, specify index")
	flag.IntVar(&fl.Complete, "complete", -1, "Complete a todo, specify index")
	flag.BoolVar(&fl.List, "list", false, "List all todo's")

	flag.Usage = func() {
		fmt.Println("Usage of this program:")
		fmt.Println("  -list")
		fmt.Println("        List all todo items")
		fmt.Println("  -add string")
		fmt.Println("        Add todo item (default 'none')")
		fmt.Println("  -delete index")
		fmt.Println("        Delete todo item on index (default -1)")
		fmt.Println("  -complete index")
		fmt.Println("        Completion status of the todo item on index (default -1)")
	}

	flag.Parse()
	fl.Execute(todos)
}

func (fl *Flags) Execute(todos *Todos) {
	todo := *todos

	switch {
	case fl.List:
		displayTodos(todo)
	case fl.Add != "":
		todos.add(fl.Add)
	case fl.Delete >= 0:
		todos.delete(fl.Delete)
	case fl.Complete >= 0:
		todos.finish(fl.Complete)
	default:
		err := errors.New("invalid flag used, try -help")
		fmt.Println(err)
	}
}
