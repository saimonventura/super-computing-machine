// For continued
// The init and post statements are optional.

package main

import "fmt"

func main() {
	sum := 1
	// NOTE: this is a way to declare and initialize a variable
	// NOTE: the values range is [0, 1000) 💨
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
