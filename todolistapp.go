package main //indicates file belongs to an executable program

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Goland Full Course"
	var rewardDessert = "Reward myself with a donut"
	var maxItems = 20
	// array is fixed size put a numberin the brackets, if u put more than the number it goes out of index
	//slice is dynamic array
	var taskItem = []string{shortGolang, fullGolang, rewardDessert}
	fmt.Println(("##### Welcome to our Todolist App! #####"))
	printTask(taskItem, maxItems)
	fmt.Println()
	//update list with new tasks
	taskItem = addTask(taskItem, "Go for a run")
	taskItem = addTask(taskItem, "Practice coding in Go")

	fmt.Println()
	printTask(taskItem, 20)
}
func printTask(taskItem []string, itemLimit int) {
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
	//printTask(updatedTaskItems, 20)
	return updatedTaskItems
}
