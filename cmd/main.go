package main

import (
	"fmt"
)

func main() {
	// if !checkFileExists("../data/todos.json") {
	// 	os.MkdirAll("../data/", 0700)
	// 	os.NewFile("../data/",)
	// }
	data, err := readJSON()
	// data, err := readJSON("$HOME/.config/godo/todo.json")
	if err != nil {
		fmt.Println("Failed to read JSON: ", err)
		return
	}
	newFlags(&data)
	// getFlagData(data)

	// displayTodos(data)
}

func displayTodos(todo Todos) {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Cyan = "\033[36m"
	var Yellow = "\033[33m"

	fmt.Println(Yellow + "Todo list:" + Reset)
	for i := range todo {
		status := " "
		if todo[i].Completed {
			status = Green + "✓" + Reset
		} else {
			status = Red + "✗" + Reset
		}
		fmt.Printf(Cyan+"[%d] "+Reset+"%s [%s]\n", i, todo[i].Title, status)
	}
}

// func getFlagData(todo Todos) {

// 	flag.Usage = func() {
// 		fmt.Println("Usage of this program:")
// 		fmt.Println("  -t string")
// 		fmt.Println("        Title of the todo item (default \"Default Title\")")
// 		fmt.Println("  -c")
// 		fmt.Println("        Completion status of the todo item (default false)")
// 	}

// 	flag.Parse()

// }
