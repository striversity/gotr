package main

import (
	"fmt"

	"github.com/striversity/projects/go-on-the-run/019-reflection/part-3/ex3/foo"
)

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "John Smith"})
	fmt.Println(&Person{name: "John Smith"})
}
