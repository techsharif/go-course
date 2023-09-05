package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("============== TODO LIST ===============")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Println("========================================")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName := scanner.Text()
			todoList.AddTask(taskName)
		case "2":
			todoList.ViewTask()
		case "3":
			fmt.Print("Enter task ID: ")
			scanner.Scan()
			taskIdStr := scanner.Text()
			taskId, err := strconv.Atoi(taskIdStr)
			if err != nil {
				fmt.Println("Invalid task id")
				continue
			}
			todoList.MarkTaskAsDone(taskId)
		case "4":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Invalid choice. Please try again")
		}
	}
}
