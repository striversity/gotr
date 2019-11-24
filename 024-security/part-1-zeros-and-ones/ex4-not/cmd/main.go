package main

import "fmt"

func main() {
	mesg := []byte("Hello")

	for _, v := range mesg {
		fmt.Printf("%8c, ", v)
	}
	fmt.Println()

	for _, v := range mesg {
		fmt.Printf("%8b, ", v)
	}
	fmt.Println()

	// bitwise NOT operation
	for _, v := range mesg {
		fmt.Printf("%8b, ", ^v)
	}
	fmt.Println()
}
