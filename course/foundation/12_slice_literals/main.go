package main

import "fmt"

func main() {
	// Common collection type in Go
	// Dynamic and can grow
	// []T{} // T is the type of elements in the slice

	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(names, names[0], names[len(names)-1])

	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20, 30)

	fmt.Println(nums)

}
