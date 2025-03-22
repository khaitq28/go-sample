package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

var tasks []Task
var currentID = 1

func clearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func addTask() {
	var title, description string
	fmt.Print("Enter task title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter task description: ")
	fmt.Scanln(&description)

	if title == "" {
		fmt.Println("Error: Title cannot be empty!")
		return
	}

	task := Task{
		ID:          currentID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, task)
	currentID++
	fmt.Println("Task added successfully!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}
	fmt.Println("\nCurrent Tasks:")
	fmt.Println("-------------")
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\n\n",
			task.ID, task.Title, task.Description, status)
	}
}

func completeTask() {
	var id int
	fmt.Print("Enter task ID to complete: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Error: Please enter a valid number!")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func deleteTask() {
	var id int
	fmt.Print("Enter task ID to delete: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Error: Please enter a valid number!")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func main() {
	for {
		err := clearScreen()
		if err != nil {
			fmt.Println("Warning: Could not clear screen:", err)
		}

		fmt.Println("Task Management System")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("\nEnter your choice (1-5): ")
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error: Please enter a valid number!")
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			completeTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
	}
}
