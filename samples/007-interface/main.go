package main

import "fmt"

type animal interface {
	getAnimalName() string
}

type dog struct{}
type cat struct{}

func main() {
	d := dog{}
	c := cat{}

	printAnimalName(d)
	printAnimalName(c)
}

func printAnimalName(a animal) {
	fmt.Println(a.getAnimalName())
}

func (dog) getAnimalName() string {
	return "Dog"
}

func (cat) getAnimalName() string {
	return "Cat"
}
