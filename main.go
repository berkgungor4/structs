package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	berk := person{
		firstName: "Berk",
		lastName:  "Gungor",
		contactInfo: contactInfo{
			email:   "berkgngr@hotmail.com",
			zipCode: 12345,
		},
	}
	berk.updateName("B.e.r.k")
	berk.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
