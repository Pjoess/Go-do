package main

import (
	// "flag"
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
}

func init() {
	fmt.Println("init state")
}
