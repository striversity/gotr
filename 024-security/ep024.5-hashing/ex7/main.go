package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dic    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	dicLen = int64(len(dic))
)

func main() {
	var v string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v = scanner.Text()
		fmt.Println(v, hash(v), string(hash2(v)))
	}
}
func hash2(s string) []byte {
	var v []byte

	h := hash(s)

	i := 0
	for h != 0 {
		i = int(h % dicLen)
		if i < 0{
			i *= -1
		}
		
		h /= dicLen
		v = append(v, dic[i])
	}

	return v
}

func hash(s string) int64 {
	var v int64
	const batches = 8

	for i, b := range s {
		idx := i % batches

		switch idx {
		case 0:
			v += int64(b) << 56
		case 1:
			v += int64(b) << 48
		case 2:
			v += int64(b) << 40
		case 3:
			v += int64(b) << 32
		case 4:
			v += int64(b) << 24
		case 5:
			v += int64(b) << 16
		case 6:
			v += int64(b) << 8
		case 7:
			v += int64(b)
		}
	}
	return v
}
