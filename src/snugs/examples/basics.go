package examples

import "fmt"

func variables() {
	// Variable declarations
	var h int
	h = 1
	j := 1
	h = h + j

	// More useful variable declaration and assignment is := operator
	taco := "hello\n"
	fmt.Printf(taco)

	// Only the postfix operator work, not the prefix uninary incrementor ++h
	h++

	// A struct variable
	var a struct {
		First string
		Last  string
	}

	a.First = "Hello"
	a.Last = "World"

}

func aSliceAtATime() {
	// Slices, not an array.  Its important ;)

	// Making an array of any size, need to use a for to initialize it
	vals2 := make([]int, 10)
	for i := 0; i < len(vals2); i++ {
		vals2[i] = i
	}

	// statically defined initializer, this is when you know what values you want to initialize with
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vals[0]++

	// Since these are slices, you can partition them easily
	valsPre := vals[3:]
	valsPost := vals[:3]

	fmt.Printf("Pre [3:] %v", valsPre)
	fmt.Printf("Post [:3] %v", valsPost)
}

func crappy() {
	// Defer waits till after the body of the function has exited.  This is in leu of try/catch
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from: %+v\n", r)
		}
	}()

	// While it feels natural to want to panic, this is a big NO NO.
	// Panic will freeze the entire process including all goroutines, use wisely.
	panic(fmt.Errorf("Oh shiz we haz an errorz"))
}

func notCrappy() error {
	return fmt.Errorf("Oh shiz we haz an errorz")
}

func selectAllTheThings() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	fmt.Printf("Channel has %v items\n", len(c))

	for {
		select {
		case c := <-c:
			fmt.Printf("Received %v\n", c)
		}

		if len(c) == 0 {
			break
		}
	}
}

func showMeTheLoops() {
	// loops
	// In most languages there are do/while, while, for, and foreach loops.
	// Go has 1 construct for looping.  for

	// Basic for loop
	for i := 0; i < 10; i++ {
		// Looping yay
	}

	// Range for loop
	var sum int
	for _ /*index*/, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		sum += v
	}

	// Range on a map
	var kk string
	m := map[string]string{"hello": "world"}
	for k, v := range m {
		kk += (k + v)
	}

}

func conditionals() {
	if true {
		fmt.Printf("Of course this is true\n")
	} else if true {
		fmt.Printf("This would be true too")
	} else {
		fmt.Print("This is just never going to be called.")
	}
}

// Basics illustrates the basic syntax of Go
func Basics() {
	fmt.Printf("----- Example 1 -----")

	// Pretty self explanatory
	variables()

	// Slices
	aSliceAtATime()

	// Conditionals
	conditionals()

	// Defer, Panic, Recover
	crappy()

	// Errors should always be returned, never panic'd
	// We will talk on this pattern more later.
	if err := notCrappy(); err != nil {
		fmt.Printf("Oops we had an error occur, we should probably handle this code: %+v\n", err)
	}

	// Looping
	showMeTheLoops()

	// select construct
	selectAllTheThings()
}
