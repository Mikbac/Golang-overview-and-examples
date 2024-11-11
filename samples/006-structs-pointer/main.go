package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email string
}

func main() {
	bob := person{"Bob", "Beaver", 22, contactInfo{"1@1.pl"}}

	alice := person{
		firstName: "Alice",
		lastName:  "Castor",
		age:       30,
		contact: contactInfo{
			email: "2@2.pl",
		}}

	var eve person
	eve.firstName = "Eve"

	bob.print()
	fmt.Println(alice)
	fmt.Println(eve.firstName)
}

func (p person) print() {
	fmt.Println(p)
}
