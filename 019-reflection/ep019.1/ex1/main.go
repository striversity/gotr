package main

import (
	"fmt"
)

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	fmt.Println(true)
	fmt.Println(2010)
	fmt.Println(9.15)
	fmt.Println(7 + 7i)
	fmt.Println("Hello world!")
	fmt.Println(ID("19520925"))
	fmt.Println([5]byte{})
	fmt.Println([]byte{})
	fmt.Println(map[string]int{})
	fmt.Println(Person{name: "Jane Doe"})
	fmt.Println(&Person{name: "Jane Doe"})
	fmt.Println(make(chan int))

}
