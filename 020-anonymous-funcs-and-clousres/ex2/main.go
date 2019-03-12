package main

import "fmt"

func main() {
	// call function that returns a function
	f := getWorkToDoLater()
	f() // do work
}

// getWorkToDoLater returns a function that will be called after it returns
func getWorkToDoLater() func() {
	fmt.Println("Entering - getWorkToDoLater() called")
	defer fmt.Println("Leaving - getWorkToDoLater()")

	doWork := func() {
		fmt.Println("doWork() called")
	}

	return doWork
}
