package examples

import (
	"fmt"
)

// BetterPerson defines a better person than person
type BetterPerson struct {
	// We are composing BetterPerson of	Person
	Person
}

// Nameable describes
type Nameable interface {
	ToName() string
}

// ToName returns a formatted name
func (bp BetterPerson) ToName() string {
	return fmt.Sprintf("%s %s", bp.First, bp.Last)
}

// typing will show how 2 types have the exact same shape and are assignable because of it.
func typing() {
	var person struct {
		First string
		Last  string
	}

	var kush struct {
		First string
		Last  string
	}

	person = Person{First: "Taco", Last: "Vendor"}

	// Illustration of Static Duct Typing
	kush = person

	fmt.Printf("What is kush? %+v\n", kush)

	// BetterPerson isn't a Person, it just behaves like one.
	var other BetterPerson
	other.Person = kush

	fmt.Printf("Are we a better Person? %+v\n", other)
	fmt.Printf("Does %s have the last name %s?\n", other.First, other.Last)
	fmt.Printf("A formatted name %s", other.ToName())

	op := func(p Nameable) {
		fmt.Printf("This should also work, %+v\n", p.ToName())
	}

	op(&other)

}

func typeSwitch() {
	findMe := func(i interface{}) {
		switch v := i.(type) {
		case string:
			fmt.Printf("Found a string: %+v\n", v)
		case Person:
			fmt.Printf("Found a person: %+v\n", v)
		}
	}

	findMe("hello")
	findMe(Person{First: "Dwayne", Last: "Johnson"})
}

// Example2 covers structs
func Example2() {
	fmt.Println("----- Example 2 -----")
	typing()
	typeSwitch()
}
