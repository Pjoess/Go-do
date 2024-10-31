package main

import (
	"encoding/json"
	"os"
)

func readJSON(filename string) (Todos, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos Todos
	err = json.Unmarshal(data, &todos)
	return todos, err

}

func writeJSON(filename string, todos Todos) error {
	// add todo to the json file
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// func checkFileExists(filename string) bool {
// 	_, error := os.Stat(filePath)
// 	//return !os.IsNotExist(err)
// 	return !errors.Is(error, os.ErrNotExist)
// }
