package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}

func deleteNote(fileName string, noteId int) {
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
		if noteId != count {
			tempFile.WriteString(sc.Text() + "\n")
		}
	}
	os.Remove(fileName)
	os.Rename("temp.txt", fileName)
}

func main() {
	var fileName string = "notes.txt"
	createFile(fileName)
	
	runApp := true
	for runApp == true {
		fmt.Print("todo: ")
		reader := bufio.NewReader(os.Stdin)
		var userCommand string = ""
		fmt.Scanf("%s", &userCommand)

		if userCommand == "exit" {
			fmt.Println("Bye")
			return
		} else if userCommand == "help" {
			help()
		} else if userCommand == "clear" {
			fmt.Print("\033[H\033[2J")
		} else if userCommand == "add" {
			fmt.Print("enter a note: ")
			note, _ := reader.ReadString('\n')
			addNote(fileName, note)
		} else if userCommand == "list" {
			listNotes(fileName)
		} else if userCommand == "delete" {
			fmt.Print("enter note id: ")
			var noteId int = 0
			fmt.Scanf("%d", &noteId)
			deleteNote(fileName, noteId)
		} else if len(userCommand) >= 1 {
			fmt.Printf("command not found: %s\n", userCommand)
		} 
	}
}