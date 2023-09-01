package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    // NOTE: If we remove the deferred function from f 
    // the panic is not recovered and reaches the top of the goroutine’s call stack, 
    // terminating the program
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

// This Go code demonstrates the use of the defer, panic, and recover statements in Go. 
// The main function calls the f function, which in turn calls the g function.

// Inside the g function, an if statement is used to check if the value of i is greater than 3. 
// If it is, the panic function is called with a formatted string that includes the value of i. 
// If i is less than or equal to 3, the defer statement is used to print a message to the console, 
// and then the g function is called recursively with i + 1.

// Inside the f function, a defer statement is used to define a function literal that calls the recover function. 
// If a panic occurs inside the g function, the recover function will return the value that was passed to the panic function. 
// The if statement inside the function literal checks if the value returned by recover is not nil. 
// If it is not nil, a message is printed to the console that indicates that 
// a panic was recovered and includes the value that was passed to the panic function.

// When the g function is called with an argument of 0, it will panic when i is incremented to 4. 
// This will cause the recover function inside the defer statement in the f function to be called, 
// which will recover from the panic and print a message to the console indicating that a panic was recovered and including the value 4.

// After the f function returns, the message "Returned normally from f." is printed to the console.

// In summary, this code demonstrates how to use the defer, panic, and recover statements in Go to handle panics and recover from them.