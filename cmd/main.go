package main

import (
	"fmt"
)

func main() {
	// if !checkFileExists("../data/todos.json") {
	// 	os.MkdirAll("../data/", 0700)
	// 	os.NewFile("../data/",)
	// }
	data, err := readJSON("../data/todos.json")
	if err != nil {
		fmt.Println("Failed to read JSON: ", err)
		return
	}
	newFlags(&data)

	// getFlagData(data)

	// displayTodos(data)
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

// func getFlagData(todo Todos) {

// 	flag.Parse()

// }
