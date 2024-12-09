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
	taskCarrots := "Buy 3 carrots"

	fmt.Println()
	fmt.Println("All tasks:")
	fmt.Println("- " + taskDishes)
	fmt.Println("- " + taskLaundry)
	fmt.Println("- " + taskCarrots)

	fmt.Println()
	fmt.Println("Priority:")
	fmt.Println("- " + taskDishes)

	fmt.Println()
	fmt.Println("Home work:")
	fmt.Println("- " + taskDishes)
	fmt.Println("- " + taskLaundry)

	fmt.Println()
	fmt.Println("Groceries:")
	fmt.Println("- " + taskCarrots)

}
