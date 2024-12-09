package main

import (
	"fmt"
)


func main() {
	fmt.Println("#######################")
	fmt.Println("#### My TODO List! ####")
	fmt.Println("#######################")
	fmt.Println()

	var taskDishes = "Do the dishes"
	var taskLaundry = "Finish the laundry"
	var taskCarrots = "Buy 3 carrots"
	
	var taskItems = []string {taskDishes, taskLaundry, taskCarrots}

	showTasks(taskItems)
	taskItems = addNewTask(taskItems, "Learn my Hindi lessons")
	taskItems = addNewTask(taskItems, "Go to gym")
	showTasks(taskItems)
}

func showTasks(taskItems []string) {
	fmt.Println("All tasks:")
	for index, task := range taskItems {
		fmt.Printf("%d) %s\n", index + 1, task)
	}
	fmt.Println()
}

func addNewTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
