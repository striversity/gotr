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
	fmt.Printf("type is '%T', value: %v\n", x, x)
}
func Println2(x int) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}
