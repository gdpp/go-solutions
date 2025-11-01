package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	gus := person{firstName: "Gustavo", lastName: "Perez"}
	fmt.Println(gus)

	var leslie person

	leslie.firstName = "Leslie"
	leslie.lastName = "Rivera"

	fmt.Println(leslie)
	fmt.Printf("%+v", leslie)

	faker := person{
		firstName: "Faker",
		lastName:  "T1",
		contactInfo: contactInfo{
			email:   "faker@t1.com",
			zipCode: 44530,
		},
	}

	// fakerPointer := &faker
	faker.updateName("Chovy")
	faker.printPerson()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
