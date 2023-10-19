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

	mySlice := []string{"Hey", "There", "How", "Are", "You"}

	updateSLice(mySlice)

	fmt.Println(mySlice)

	alex.updateName("Sandor")

	alex.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateSLice(s []string) {
	s[0] = "Bye"
}

// Turn address into a value with *address
// Turn value into a address with &value
