package main

import "github.com/striversity/projects/go-on-the-run/019-reflection/part-3/ex1/foo"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	foo.Println(true)
	foo.Println(2010)
	foo.Println(9.15)
	foo.Println(7 + 7i)
	foo.Println("Hello world!")
	foo.Println(ID("19520925"))
	foo.Println([5]byte{})
	foo.Println([]byte{})
	foo.Println(map[string]int{})
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "Jane Doe"})
	foo.Println(make(chan int))
}
