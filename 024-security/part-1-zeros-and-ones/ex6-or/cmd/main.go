package main

import "fmt"

func main() {
	mesg := []byte("Hello")

	fmt.Printf("%10s", "Char:")
	for _, v := range mesg {
		fmt.Printf("%8c, ", v)
	}
	fmt.Println()
	
	fmt.Printf("%10s", "Binary:")
	for _, v := range mesg {
		fmt.Printf("%8b, ", v)
	}
	fmt.Println()
	
	// bitwise OR operation
	k := byte('V')
	
	fmt.Printf("%10s", "k-value:")
	for range mesg {
		fmt.Printf("%8b, ", k)
	}
	fmt.Println()
	
	fmt.Printf("%10s", "Result:")
	for _, v := range mesg {
		fmt.Printf("%8b, ", k|v)
	}
	fmt.Println()
}
