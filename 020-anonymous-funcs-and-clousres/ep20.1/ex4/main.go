package main

import "fmt"

func main() {
	// call function that returns a function
	f := getWorkToDoLater()
	// very long time ...
	x := f(4) // do work
	fmt.Printf("f(4) = %v\n", x)
}

// getWorkToDoLater returns a function that will be called after it returns
func getWorkToDoLater() func(int) int {
	fmt.Println("Entering - getWorkToDoLater() called")
	defer fmt.Println("Leaving - getWorkToDoLater()")
	var b = 11 // b.ref = 1

	doWork := func(x int) int {
		fmt.Printf("doWork() called to double %v plus %v\n", x, b)
		return 2*x + b // b.ref = 2
	}
	return doWork
} // b.ref = 1
