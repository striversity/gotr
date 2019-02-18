package main

import "fmt"

func main() {
	var x interface{}
	x = 3.14
	fmt.Printf("t: %T, v: %v\n", x, x)
	goo := x
	fmt.Printf("goo: t: %T, v: %v\n", goo, goo)
	x = &struct{ name string }{}
	fmt.Printf("t: %T, v: %v\n", x, x)
	hoo := x
	fmt.Printf("hoo: t: %T, v: %v\n", hoo, hoo)
}
