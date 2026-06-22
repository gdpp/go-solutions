package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(&a, &b)

	fmt.Println(a, b) // Should print 20, 10
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
