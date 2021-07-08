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
	joe := person{firstName: "Joe", lastName: "Collins"}
	fmt.Println(joe)

	var tom person
	fmt.Printf("%+v\n", tom)
	tom.firstName = "Tom"
	fmt.Printf("%+v\n", tom)

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 01776,
		},
	}
	alex.print()
	alex.updateName("Alexander")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p &person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
