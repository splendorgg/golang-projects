package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID       int
	Title    string
	Complete bool
}

var tasks []Task

func addTask(title string) {
	newTask := Task{
		ID:       len(tasks) + 1,
		Title:    title,
		Complete: false,
	}
	tasks = append(tasks, newTask)
}

func listTask() {
	fmt.Println("To-Do List:")
	for _, task := range tasks {
		status := " "
		if task.Complete == true {
			status = "X"
		}
		fmt.Printf("Status:[%s] Task Number:%d. Task:%s\n", status, task.ID, task.Title)
	}
}

func completeTask(title string) {
	taskFound := false
	pntTask := &taskFound
	for i := range tasks {
		if title == tasks[i].Title {
			tasks[i].Complete = true
			fmt.Println("This task completed")
			*pntTask = true
			break
		}
	}
	if *pntTask == false {
		fmt.Printf("Task '%s' not found\n", title)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter 'add','list','complete'or 'exit': ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "add":
			fmt.Print("Add a task: ")
			scanner.Scan()
			title := scanner.Text()
			addTask(title)
			fmt.Println("Your task successfully added to the list")
		case "list":
			listTask()
		case "complete":
			fmt.Print("Which task do you want to complete?: ")
			scanner.Scan()
			title := scanner.Text()
			completeTask(title)
		case "exit":
			fmt.Println("Exited")
			return
		default:
			fmt.Println("Invalid choice. Please enter 'add','list' or 'exit'")
		}
	}
}
