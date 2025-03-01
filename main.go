package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Println("============ Todo ============")
	fmt.Println("App version 1.0")
	fmt.Println("Author: Alex Rodriguez")
	fmt.Println()
	fmt.Println("Commands: ")
	fmt.Println("help             - displays all commands")
	fmt.Println("add              - adds a note")
	fmt.Println("list             - lists all notes")
	fmt.Println("delete           - deletes a note by id, id must be an integer")
	fmt.Println("clear            - clears the terminal")
	fmt.Println("exit             - exits the app")
	fmt.Println()
}

// Creates file if it doesn't exist
func createFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		os.Create(fileName)
	}
	defer file.Close()
}

func addNote(fileName string, n string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("note added successfully!")
}

func listNotes(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var noteId int = 0
	for sc.Scan() {
		noteId += 1
		fmt.Printf("%d. %s\n", noteId, sc.Text())
	}
	if noteId == 0 {
		fmt.Println("You don't have any notes...")
		return
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}

func deleteNote(fileName string, noteId string) {
	id, err := strconv.Atoi(noteId)
	if err != nil {
		fmt.Println("error: please enter a number greater than 1 for the note id")
		return
	}
	if id <= 0 {
		fmt.Println("error: please enter a number greater than 1 for the note id")
		return
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.Create("temp.txt")
	tempFile, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	var count int = 0
	for sc.Scan() {
		count += 1
		if id != count {
			tempFile.WriteString(sc.Text() + "\n")
		}
	}
	os.Remove(fileName)
	os.Rename("temp.txt", fileName)
	fmt.Printf("note with id %d was deleted successfully!\n", id)
}

func parseUserInput(c string) string{
	reader := bufio.NewReader(os.Stdin)
	userCommandLine, _ := reader.ReadString('\n')
	var firstWord string = ""
	if len(strings.Fields(userCommandLine)) > 0 {
		firstWord = strings.Fields(userCommandLine)[0]
	}

	switch c {
	case "addNote":
		return userCommandLine // Returns the whole user command line, for new notes
	default:
		return firstWord // Returns the first word in the command line, for getting commands or ids
	}
}

func main() {
	var fileName string = "notes.txt"
	createFile(fileName)

	runApp := true
	for runApp == true {
		fmt.Print("todo: ")
		command := parseUserInput("command")

		switch command {
		case "add":
			fmt.Print("enter a note: ")
			addNote(fileName, parseUserInput("addNote"))
		case "list":
			listNotes(fileName)
		case "delete":
			fmt.Print("enter note id: ")
			deleteNote(fileName, parseUserInput("deleteNote"))
		case "help":
			help()
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "exit":
			fmt.Println("Bye")
			return
		case "":
		default:
			fmt.Printf("command not found: %s\n", command)
		}
	}
}
