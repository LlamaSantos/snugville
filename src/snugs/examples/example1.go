/*
Example 1:
Illustrates some basic fmt package uses
*/

package examples

import (
	"fmt"
)

// Example1 is a simple hello world example using the fmt package
func Example1() {
	fmt.Println("----- Example 1 -----")

	// Printf doesn't add a newline
	fmt.Printf("Hello World\n")

	// Println Does
	fmt.Println("Hello World")

	// Println can't do this tho
	fmt.Printf("Hello world for the %vrd time\n", 3)

	fmt.Printf("Hello from a Person %+v\n", Person{First: "Hello", Last: "World"})

	err := fmt.Errorf("Creates a new error with %s", "this arg")
	fmt.Printf(err.Error())
}
