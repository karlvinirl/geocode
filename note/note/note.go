package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string
	Content string
	CreatedAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input for note constructor")
	}
	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your not titled: %v has the following content:\n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}