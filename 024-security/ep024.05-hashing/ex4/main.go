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
	for i := 0; i < len(s)/4; i++ {
		ss := s[:4]
		s = s[4:]

		v += int64(ss[0])<<24 + int64(ss[1])<<16 + int64(ss[2])<<8 + int64(ss[3])
	}
	return v
}
