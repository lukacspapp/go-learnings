package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {

	alexContact := contact{email: "test@mail.com", zip: 1045}

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   alexContact,
	}

	alex.updateName("Sandor")

	alex.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Turn address into a value with *address
// Turn value into a address with &value
