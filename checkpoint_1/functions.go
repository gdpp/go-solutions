package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello " + name)
}

func sum(a int, b int) int {
	return a + b
}

func average(nums []int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	average := total / len(nums)

	return average
}

func countDown(counter int) {
	for counter > 0 {
		fmt.Println("Counter: ", counter)
		counter = counter - 1
	}
}

func showItems(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func evenOrOdd(items []int) {
	for _, n := range items {
		if n%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
