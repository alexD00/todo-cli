package main

import (
	"fmt"
)

func help() {
	fmt.Println("============ Todo ============")
	fmt.Println("App version 1.0")
	fmt.Println("Author: Alex Rodriguez")
	fmt.Println()
	fmt.Println("Commands: ")
	fmt.Println("help             - displays all commands")
	fmt.Println("add <note>       - adds a note")
	fmt.Println("list             - lists all notes")
	fmt.Println("delete <note id> - deletes a note by id")
	fmt.Println("clear            - clears the terminal")
	fmt.Println("exit             - exits the app")
	fmt.Println()
}

func main() {
	runApp := true
	for runApp == true {
		var userCommand string = ""
		fmt.Print("todo: ")
		fmt.Scanf("%s", &userCommand)

		if userCommand == "exit" {
			fmt.Println("Bye")
			return
		} else if userCommand == "help" {
			help()
		} else if userCommand == "clear" {
			fmt.Print("\033[H\033[2J")
		} else if len(userCommand) >= 1 {
			fmt.Printf("command not found: %s\n", userCommand)
		}
	}
}