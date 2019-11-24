package main

import "fmt"

func main() {
	mesg := "012345"

	for _, v := range mesg {
		fmt.Printf("c: %v, %b", v, v)
	}
	fmt.Println()
}
