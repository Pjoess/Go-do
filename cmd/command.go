package main

import "flag"

type Flags struct {
	Add      string
	Delete   int
	Complete int
	List     bool
	Nil      bool
}

func newFlags(){
	fl := Flags{}

	flag.StringVar(&fl.Add, "add", "", "Add a new todo, specify title")
	flag.IntVar(&fl.Delete, "delete", -1, "Remove a todo, specify index")
	flag.IntVar(&fl.Complete, "complete", -1, "Complete a todo, specify index")
	flag.BoolVar(&fl.List, "add", false, "List all todo's")

	flag.Parse()
}

func (fl *Flags) Execute(todos *Todos) {
	todo := *todos

	switch {
	case fl.Nil:
		displayTodos(todo)
	case fl.List:
		displayTodos(todo)
	}
}
