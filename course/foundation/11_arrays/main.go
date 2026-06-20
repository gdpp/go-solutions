package main

import "fmt"

func main() {
	// Fixed and cannot grow -> vvi
	var marks [3]int
	marks[0] = 90
	marks[1] = 80
	marks[2] = 70

	// Array literal
	scores := [5]int{100, 90, 80, 70, 60}

	fmt.Println(len(scores))
}
