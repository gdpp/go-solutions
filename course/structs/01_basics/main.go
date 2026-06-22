package main

import "fmt"

// Struct groups related fields into one type
type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		Id:    1,
		Name:  "Gustavo",
		Email: "gdpp@dev.com",
		Age:   33,
	}

	fmt.Println(u1, u1.Name, u1.Age)

	u2 := User{
		Id: 2,
	}

	fmt.Println("Partial user: ", u2)
}
