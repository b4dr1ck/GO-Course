package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"example.com/note/note"
)

func main() {
	listDir()

	fmt.Println("Add a new note:")

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Output()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving file")
		return
	}
}

func listDir() {
	files, err := os.ReadDir(".")

	fmt.Println("Your saved notes:")

	if err != nil {
		fmt.Println("Error reading directory")
		return
	}

	var fileName string
	var fileContent []byte
	for _, file := range files {
		matched, _ := regexp.MatchString(`.*\.json$`, file.Name())
		if matched {
			fileName = file.Name()
			fileContent, err = os.ReadFile(file.Name())
			if err != nil {
				fmt.Printf("Error reading file %v\n", fileName)
				continue
			}

			var note note.Note
			err = json.Unmarshal(fileContent, &note)
			if err != nil {
				fmt.Printf("Error parsing file %v\n", fileName)
				continue
			}
			note.Display()
			fmt.Println()
		}
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// using IO-Buffer for input to be able to use multiple words
	// fmt.Scan() can only handle single words!
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows line-endings

	return text
}
