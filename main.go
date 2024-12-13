package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Do the laundry", "Clean the house", "Buy groceries"}

func main() {
	var port string = "8080"
	fmt.Printf("Running the TODO List App on http://localhost:%s ...\n", port)

	http.HandleFunc("/", welcomePage)
	http.HandleFunc("/list_tasks", listTasks)

	http.ListenAndServe(":" + port, nil)
}

func welcomePage(writer http.ResponseWriter, request *http.Request) {
	var message = "Hello, welcome to the TODO List App!"
	fmt.Fprintln(writer, message)
}

func listTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
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
