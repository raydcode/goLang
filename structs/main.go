package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
}

// Greeting Method

func (p Person) greet() string {
	return "Hello myself" + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
	// Init Person using struct

	person := Person{
		firstName: "kathir",
		lastName:  "karthik",
		gender:    "male",
		age:       21,
	}

	// Shorthand

	otherPerson := Person{
		"mithila",
		"palkar",
		"female",
		25,
	}

	fmt.Println(person)
	fmt.Println(otherPerson)

	fmt.Println(person.firstName, "&", otherPerson.firstName)

	fmt.Println(person.greet())
}
