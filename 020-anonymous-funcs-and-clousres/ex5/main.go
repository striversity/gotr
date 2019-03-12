package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func main() {
	// call function that returns a function
	f := getWorkToDoLater()
	// very long time ...
	x := f(4) // do work
	fmt.Printf("f(4) = %v\n", x)
}

// getWorkToDoLater returns a function that will be called after it returns
func getWorkToDoLater() func(uint8) *Person {
	fmt.Println("Entering - getWorkToDoLater() called")
	defer fmt.Println("Leaving - getWorkToDoLater()")
	var setAge func(uint8) *Person
	var b = &Person{Name: "John"}
	{
		setAge = func(age uint8) *Person {
			fmt.Printf("Updating %v's age to %v\n", b.Name, age)
			b.Age = age
			return b
		}
	}
	return setAge
}
