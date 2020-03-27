package main

import "fmt"

func main() {
	names := []string{"Guyana", "United States", "Canada"}

	for _, v := range names {
		fmt.Println(v, hash(v))
	}
}

func hash(s string) int {
	v := 0
	for _, b := range s {
		v += int(b)  // v = v + int(b)
	}
	return v
}
