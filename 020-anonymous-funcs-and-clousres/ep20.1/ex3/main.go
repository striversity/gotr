package main

import "fmt"

func main() {
	// call function that returns a function
	f := getWorkToDoLater()
	x := f(4) // do work
	fmt.Printf("result of work: %v\n", x)
}

// getWorkToDoLater returns a function that will be called after it returns
func getWorkToDoLater() func(int) int {
	fmt.Println("Entering - getWorkToDoLater() called")
	defer fmt.Println("Leaving - getWorkToDoLater()")

	doWork := func(x int) int {
		fmt.Printf("doWork() called to double %v\n", x)
		return 2 * x
	}
	return doWork
}
