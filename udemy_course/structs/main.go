package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex)

	// var sky person

	// fmt.Printf("%+v\n", sky)

	// sky.firstName = "Skylord"
	// sky.lastName = "Perry"

	// fmt.Printf("%+v\n", sky)

	john := person{
		firstName: "John",
		lastName: "Doe",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// johnPointer := &john
	// johnPointer.updateName("johnny")
	john.updateName("johnny")
	john.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}