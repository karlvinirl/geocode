package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}


func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Printf("Saving data failed")
		return err
	}
	return nil
	
}

func outputAndSave(data outputable) error {
	err := saveData(data)
	if err != nil {
		return err
	}
	data.Display()
	return nil
}

func main() {

	title, content := getNoteData()
	myNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputAndSave(myNote)
	if err != nil {
		return
	}

	content = getTodoData()
	myTodo, err := todo.New(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputAndSave(myTodo)
	if err != nil {
		return
	}
}


func getTodoData() string {
	content := getUserInput("Todo content:")
	return content
}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
