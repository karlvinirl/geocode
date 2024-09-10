package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string
}	

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input for todo constructor")
	}
	return Todo {
		Content: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Your todo has the following content:\n\n%v\n", todo.Content)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}