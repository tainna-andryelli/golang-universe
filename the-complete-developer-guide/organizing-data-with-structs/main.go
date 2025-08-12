package main

import "fmt"

func main() {
	
	persona := person{
		firstName: "Tai", 
		lastName: "Ribeiro",
	}
	persona.firstName = "Luann"
	persona.lastName = "Muryell"
	persona.contact = contactInfo{
		email: "lalala@gmail.com",
		zipCode: 9292992,
	}

	persona.print()
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}