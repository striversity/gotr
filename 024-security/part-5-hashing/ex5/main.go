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
	const batches = 4

	for i, b := range s {
		idx := i % batches

		switch idx {
		case 0:
			v += int(b) << 24
		case 1:
			v += int(b) << 16
		case 2:
			v += int(b) << 8
		case 3:
			v += int(b)
		}
	}
	return v
}
