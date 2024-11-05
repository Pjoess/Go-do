package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getFilepath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home folder: ", err)
		return "", err
	}
	dir := filepath.Join(home, ".config", "godo", "todo.json")
	return dir, nil
}

func readJSON() (Todos, error) {
	dir, err := getFilepath()
	if err != nil {
		return nil, err
	}

	data, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	bytedata, _ := io.ReadAll(data)
	var todos Todos
	err = json.Unmarshal(bytedata, &todos)
	return todos, err

}

func writeJSON(todos Todos) error {
	// add todo to the json file
	dir, err := getFilepath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	file, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// func checkFileExists(filename string) bool {
// 	_, error := os.Stat(filename)
// 	//return !os.IsNotExist(err)
// 	return !errors.Is(error, os.ErrNotExist)
// }
