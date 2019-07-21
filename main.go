package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	age int
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

func main() {
	firstPerson := person{"Name01", "LastName01", 20}
	secondPerson := person{"Name02", "LastName02", 30}

	fmt.Println(firstPerson.fullName())
	fmt.Println(secondPerson.fullName())
}