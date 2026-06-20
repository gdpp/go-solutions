package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

func divide(a int, b int) (x int) {
	x = a / b

	return
}

// variadic function
func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums {
		total = total + currentValue
	}

	return total
}

func main() {
	sum1 := add(3, 4)
	sum2 := add(10, 20)

	fmt.Println(sum1)
	fmt.Println(sum2)

	s, p := sumAndProduct(3, 4)
	fmt.Println(s, p)

	res := func(n int) int {
		return n * n
	}

	fmt.Println(res(5))

	// IIFE - Immediately Invoked Function Expression
	iifeRes := func(a int, b int) int {
		return a + b
	}(5, 10)

	fmt.Println(iifeRes)
}
