package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	Println(true)
	Println(2010)
	Println(9.15)
	Println(7 + 7i)
	Println("Hello world!")
	Println(ID("19520925"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "Jane Doe"})
	Println(&Person{name: "Jane Doe"})
	Println(make(chan int))

}

// Println is my simple println function
func Println(x interface{}) {
	// x.type - for type information
	// x.value - for the value assigned
	if v, ok := x.(ID); ok {
		fmt.Printf("x has type ID, which I defined. The value is: %v\n", v)
	} else {
		fmt.Printf("'%T' is not the type I want\n", x)
	}
}
