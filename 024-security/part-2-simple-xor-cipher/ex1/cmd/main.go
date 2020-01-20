package main

import "fmt"

func main() {
	mesg := []byte("Hello")
	k := byte('a')

	fmt.Printf("Message: %s\n", string(mesg))
	enc := xor(mesg, k)
	fmt.Printf("Encypted: %s\n", string(enc))

}

func xor(buf []byte, k byte) []byte {
	out := make([]byte, len(buf))

	for i, v := range buf {
		out[i] = k ^ v
	}

	return out
}
