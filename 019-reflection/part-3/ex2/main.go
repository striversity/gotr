package main

import "github.com/striversity/projects/go-on-the-run/019-reflection/part-3/ex2/foo"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "John Smith"})
}
