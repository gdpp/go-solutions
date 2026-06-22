package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Gus", Age: 34}

	u.Birthday()

	fmt.Println("after:", u.Age)

}

// Method pointer receiver
func (u *User) Birthday() {
	u.Age++
}
