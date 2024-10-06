package todo

import (
	"encoding/json"
	"os"
)

const storageFile = "todos.json"

func (tl *TodoList) SaveTodo() bool {
	jsonString, _ := json.Marshal(tl)
	os.WriteFile(storageFile, jsonString, os.ModePerm)
	tl.LoadTodos()
	return true
}

func (tl *TodoList) LoadTodos() bool {
	data, _ := os.ReadFile(storageFile)
	json.Unmarshal(data, tl)
	return true
}
