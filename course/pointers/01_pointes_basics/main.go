package main

import "fmt"

func main() {
	// Store the memory address of any value

	// &x -> Address of x (makes a pointer)
	// *p -> deref (go to that address and read/write)

	score := 95

	fmt.Println("Before:", score)

	addScore(&score)

	fmt.Println("After:", score)
}

func addScore(score *int) {
	*score = *score + 5
}
