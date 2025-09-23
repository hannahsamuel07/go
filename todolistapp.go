package main //indicates file belongs to an executable program

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Goland Full Course"
	var rewardDessert = "Reward myself with a donut"
	// array is fixed size put a numberin the brackets, if u put more than the number it goes out of index
	//slice is dynamic array
	taskItem := []string{shortGolang, fullGolang, rewardDessert}
	fmt.Println(("##### Welcome to our Todolist App! #####"))

	fmt.Println("List of my Todos")
	fmt.Println("Tasks:", taskItem)

}
