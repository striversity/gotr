package main

import "fmt"

func main() {
	mesg := []byte("This is a very secret message!")
	k := byte('V')

	fmt.Printf("Message: %s\n", string(mesg))
	enc := xor(mesg, k)
	fmt.Printf("Encypted: %s\n", string(enc))

	// -----
	enc2 := xor(enc, k)
	fmt.Printf("Encypted: %s\n", string(enc2))

}

func xor(buf []byte, k byte) []byte {
	out := make([]byte, len(buf))

	for i, v := range buf {
		out[i] = k ^ v
	}

	return out
}
