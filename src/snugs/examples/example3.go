package examples

import "fmt"
import "sync"

// FunctionalPerson defines a person with more functionality
type FunctionalPerson struct {
	Person
	called int
	mutex  sync.Mutex
}

// ToName returns the name of a FunctionalPerson
func (fp *FunctionalPerson) ToName() string {
	// This will execute after the return statement
	defer fp.mutex.Unlock()

	// Assure that called can be updated one at a time per instance
	fp.mutex.Lock()
	fp.called++

	return fmt.Sprintf("%s, %s", fp.Last, fp.First)
}

// Called returns the number of times ToName has been called
func (fp *FunctionalPerson) Called() int {
	defer fp.mutex.Unlock()
	fp.mutex.Lock()
	return fp.called
}

// Example3 contains examples for types and receivers
func Example3() {
	fmt.Println("----- Example 3 -----")

	write := func(n Nameable) {
		fmt.Printf("Person: %s\n", n.ToName())
	}

	called := func(c Callable) {
		fmt.Printf("%+v has been called %v times\n", c, c.Called())
	}

	// Because we need to keep state, we have to initialize a new instance of FunctionalPerson.
	// The pointer receivers above on Called  and ToName enforce this.
	// Without the new operator the code below will not work.
	// `new` is a builtin function to create a new instance of a struct.
	person1 := new(FunctionalPerson)
	person1.First = "Jimmy"
	person1.Last = "Hoffa"
	write(person1)
	called(person1)
}