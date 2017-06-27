package examples

// Person describes a person
type Person struct {
	First string
	Last  string
}

// Nameable describes a way to name something
type Nameable interface {
	ToName() string
}

// Callable describes a way to get how many times something's been called
type Callable interface {
	Called() int
}
