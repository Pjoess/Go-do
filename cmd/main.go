package main

import (
	// "flag"
	// "errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	todos := Todos{}
	todos.add("hello todo")
	todos.add("delete todo")
	fmt.Printf("%+v\n", todos)
	todos.delete(1)
	fmt.Printf("%+v\n", todos)

	err := writeJSON("../data/todos.json", todos)
	if err != nil {
		fmt.Println("Failed to write JSON: ", err)
	}

	data, err := readJSON("../data/todos.json")
	if err != nil {
		fmt.Println("Failed to read JSON: ", err)
	}

	fmt.Printf("%+v\n", data)
}

func init() {
	fmt.Println("init state")
}
