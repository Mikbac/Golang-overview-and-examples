package main

import "fmt"

type person struct {
	firstName string
}

func main() {
	bob0 := person{"Bob_0"}
	bob1 := person{"Bob_1"}
	bob2 := person{"Bob_2"}

	bob0.print()        // {Bob_0}
	bobPointer := &bob0 // pointer
	bobPointer.updateFirstNameViaPointer("Booooob_0_1")
	bob0.print() // {Booooob_0_1}

	bob1.print() // {Bob_1}
	bob1.updateFirstNameViaPointer("Booooob_1_1")
	bob1.print() // {Booooob_1_1}

	bob2.print() // {Bob_2}
	bob2.updateFirstName("Booooob_2_1")
	bob2.print() // {Bob_2}

	// --------------------------------------------------------------
	fmt.Println("--------------------------------------------------------------")

	sampleSlice := []string{"aaa", "bbb", "ccc"}
	updateSlice(sampleSlice) // we copy slice, but the two slices still refer to the same array
	fmt.Println(sampleSlice) // [111 bbb ccc]
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateFirstNameViaPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// or
	// pointerToPerson.firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Println(pointerToPerson)
}

func updateSlice(s []string) {
	s[0] = "111"
}
