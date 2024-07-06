package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Simple To-Do List")
	fmt.Println("-----------------")

	// Open a file to store the to-do list items
	fileName := "todo.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Infinite loop to interact with the user
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n1. Add a task\n2. View tasks\n3. Exit\n\nChoose an option: ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print("Enter task to add: ")
			scanner.Scan()
			task := scanner.Text()

			// Write the task to the file
			_, err := fmt.Fprintln(file, task)
			if err != nil {
				fmt.Println("Error writing to file:", err)
			} else {
				fmt.Println("Task added successfully!")
			}

		case "2":
			fmt.Println("\nTasks:")
			// Read and display tasks from the file
			file.Seek(0, 0) // Move pointer to the beginning of the file
			fileScanner := bufio.NewScanner(file)
			for fileScanner.Scan() {
				fmt.Println(fileScanner.Text())
			}

		case "3":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}
