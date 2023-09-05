package main

import "fmt"

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	taskId := len(td.tasks) + 1
	task := NewTask(taskId, taskName)
	td.tasks = append(td.tasks, task)
	fmt.Println("Task Added Successfully!")
}

func (td *TodoList) ViewTask() {
	fmt.Println("=============== TASKS ===========")
	for _, task := range td.tasks {
		doneStatus := " "
		if task.Done {
			doneStatus = "x"
		}
		fmt.Printf("[%s] Task #%d: %s\n", doneStatus, task.ID, task.Name)
	}
	fmt.Println("=================================")
}

func (td *TodoList) MarkTaskAsDone(taskId int) {
	if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println("Invalid task ID.")
		return
	}
	td.tasks[taskId-1].MarkAsDone()
}
