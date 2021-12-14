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
	isMarried bool

	// Shorthand

	// firstName, lastName, gender string
	// age                         int
	// isMarried                   bool
}

// Greeting Method

func (p Person) greet() string {

	return "Hello myself " + p.firstName + "  " + p.lastName + " and I am " + strconv.Itoa(p.age) + " isMarried " + strconv.FormatBool(p.isMarried)
}

// getMarried (pointer receiver)

func (p *Person) getMarried() string {
	if !p.isMarried {
		p.isMarried = true
	}
	return p.firstName + " " + "getting Married"
}

// pointer Receiver

func (p *Person) hasBday() {
	p.age++
}

func main() {
	// Init Person using struct

	person := Person{
		firstName: "kathir",
		lastName:  "karthik",
		gender:    "male",
		age:       21,
		isMarried: false,
	}

	// Shorthand

	otherPerson := Person{
		"mithila",
		"palkar",
		"female",
		25,
		false,
	}

	fmt.Println(person)
	fmt.Println(otherPerson)

	fmt.Println(person.firstName, "&", otherPerson.firstName)

	fmt.Println(person.greet())

	person.hasBday()

	fmt.Println(person.greet())

	fmt.Println(otherPerson.getMarried())

	fmt.Println(otherPerson.greet())

}
