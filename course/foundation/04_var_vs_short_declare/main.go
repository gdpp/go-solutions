package main

import (
	"fmt"
)

func main() {
	var city string
	city = "London"

	fmt.Println("city: ", city)

	population := 40000
	population = population + 10000

	fmt.Println("Population: ", population)

	likes, comments := 10, 20
	fmt.Println("likes: ", likes, "Comments: ", comments)
}
