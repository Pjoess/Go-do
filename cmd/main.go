package main

import (
	// "flag"
	"fmt"
)

func main() {
	data, err := readJSON("../data/todos.json")
	if err != nil {
		fmt.Println("Failed to read JSON: ", err)
		return
	}

	displayTodos(data)
}

func displayTodos(todo Todos) {
	fmt.Println("Todo list:")
	for i := range todo {
		status := " "
		if todo[i].Completed {
			status = "✓"
		} else {
			status = "✗"
		}
		fmt.Printf("[%d] %s [%s]\n", i, todo[i].Title, status)
	}
}

func init() {
	fmt.Println("init state")
}
