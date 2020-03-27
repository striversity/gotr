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
	switch x.(type) {
	case bool:
		fmt.Print("This is a boolean value: ", x.(bool))
	case int:
		fmt.Print("This is my nice int value: ", x.(int))
	case float64:
		fmt.Print(x.(float64))
	case complex128:
		fmt.Print(x.(complex128))
	case string:
		fmt.Print(x.(string))
	case Person:
		fmt.Print(x.(Person))
	case chan int:
		fmt.Print(x.(chan int))
	default:
		fmt.Print("Unknown type")
	}

	fmt.Print("\n")
}
