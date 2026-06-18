package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Gustavo"
	lastName := "Perez"

	fullName := firstName + " " + lastName

	fmt.Println("Complete name: ", fullName)

	fmt.Println(strings.ToUpper(fullName))
}
