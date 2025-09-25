package main //indicates file belongs to an executable program

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Goland Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItem = []string{shortGolang, fullGolang, rewardDessert}

func main() {

	//var maxItems = 20
	// array is fixed size put a numberin the brackets, if u put more than the number it goes out of index
	//slice is dynamic array
	fmt.Println(("##### Welcome to our Todolist App! #####"))
	http.HandleFunc("/", home)

	http.HandleFunc("/show-tasks", showTasks)

	http.HandleFunc("/hello-go", helloUser)
	http.ListenAndServe(":8080", nil)
	//printTask(taskItem)
	//fmt.Println()
	//update list with new tasks
	//taskItem = addTask(taskItem, "Go for a run")
	//taskItem = addTask(taskItem, "Practice coding in Go")

	//fmt.Println()
	//printTask(taskItem)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Todo App!")
	fmt.Fprintln(w, "Try /hello-go or /show-tasks")
}
func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItem {
		fmt.Fprintln(writer, task)
	}
}
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our TodolistApp!"
	fmt.Fprintln(writer, greeting)
}

/*func printTask(taskItem []string) {
	fmt.Println("List of my Todos")
	//for loop in go, index and range keyword
	// index, -> _, //when we dont need to use the index
	for index, task := range taskItem {
		//fmt.Println(index+1, ".", task)
		fmt.Printf("%d: %s\n", index+1, task) // print f for string modication for formatting

	}

}

// inputs expected					//type of output expected
func addTask(taskItem []string, newTask string) []string {
	var updatedTaskItems = append(taskItem, newTask)
	return updatedTaskItems
}
*/
