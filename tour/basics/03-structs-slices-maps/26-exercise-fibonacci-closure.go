// Exercise: Fibonacci closure
// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// Initialize the first two values of the sequence
	a, b := 0, 1

	// Return a function that returns the next value of the sequence
	return func() int {
		// Calculate the next value of the sequence
		next := a + b

		// Update the values of the sequence
		a = b
		b = next

		// Return the next value of the sequence
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
