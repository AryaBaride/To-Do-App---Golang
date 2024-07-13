package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todoTasks = []string{}

func main() {
	fmt.Println("Welcome to To-Do App")
	for {
		fmt.Println("Press 1 to add a task")
		fmt.Println("Press 2 to delete a task")
		fmt.Println("Press 3 to view all tasks")
		fmt.Println("Press 4 to exit")

		fmt.Print("Enter your choice: ")
		input := bufio.NewReader(os.Stdin)
		choice, _ := input.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter the task you want to add: ")
			addTask, _ := input.ReadString('\n')
			addTask = strings.TrimSpace(addTask)
			todoTasks = append(todoTasks, addTask)
			fmt.Println("Task added successfully")

		case "2":
			fmt.Print("Enter the task you want to delete: ")
			deleteTask, _ := input.ReadString('\n')
			deleteTask = strings.TrimSpace(deleteTask)
			for i, task := range todoTasks {
				if task == deleteTask {
					todoTasks = append(todoTasks[:i], todoTasks[i+1:]...)
					fmt.Println("Task deleted successfully")
					break
				}
			}

		case "3":
			fmt.Println("Your tasks are:")
			for _, task := range todoTasks {
				fmt.Println(task)
			}

		case "4":
			fmt.Println("Thank you for using To-Do App")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
