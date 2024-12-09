package main

import (
	"fmt"
)

func main() {
	fmt.Println("#######################")
	fmt.Println("#### My TODO List! ####")
	fmt.Println("#######################")

	var taskDishes = "Do the dishes"
	var taskLaundry = "Finish the laundry"
	var taskCarrots = "Buy 3 carrots"

	var taskItems = []string {taskDishes, taskLaundry, taskCarrots}

	fmt.Println()
	fmt.Println("All tasks:")
	for index, task := range taskItems {
		fmt.Printf("%d) %s\n", index + 1, task)
	}

}
