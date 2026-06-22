package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Gus", Age: 34}
	fmt.Println(u.Intro())
}

// Method value receiver
// means this method receives a copy of the User
func (u User) Intro() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d", u.Name, u.Age)
}
