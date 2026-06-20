package main

import "fmt"

func main() {
	// make([]T, length, capacity)
	scores := make([]int, 0, 5)

	fmt.Println(len(scores), cap(scores))

	scores = append(scores, 100, 90, 80, 70, 60, 50, 40)

	fmt.Println(len(scores), cap(scores))

	todos := []string{"Buy groceries", "Clean the house", "Pay bills"}
	more := []string{"Go for a walk", "Read a book"}

	//...

	todos = append(todos, more...)

	fmt.Println(todos)

	views := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// for range
	total := 0

	for i, v := range views {
		fmt.Println(i, v)
		total = total + v
	}
}
