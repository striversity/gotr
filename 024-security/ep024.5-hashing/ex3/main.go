package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var v string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v = scanner.Text()
		fmt.Println(v, hash(v))
	}
}

func hash(s string) int {
	v := 0
	for _, b := range s {
		v ^= int(b) // v = v + int(b)
	}
	return v
}
