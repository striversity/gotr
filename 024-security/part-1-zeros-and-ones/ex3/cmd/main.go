package main

import "fmt"

func main() {
	mesg := "012345"

	for _, v := range mesg {
		fmt.Printf("%8c, ", v)
	}
	fmt.Println()

	for _, v := range mesg {
		fmt.Printf("%8b, ", v)
	}
	fmt.Println()
}
