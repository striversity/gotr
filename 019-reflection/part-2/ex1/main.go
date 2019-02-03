package main

import "fmt"

func main() {
	var x interface{}
	x = 3.14
	fmt.Printf("x: type = %T, value = %v\n", x, x)
	goo := x
	fmt.Printf("goo: type = %T, value = %v\n", goo, goo)
	
	x = &struct{ name string }{}
	fmt.Printf("x: type = %T, value = %v\n", x, x)
	hoo := x
	fmt.Printf("hoo: type = %T, value = %v\n", hoo, hoo)
}
