package main

import (
	"fmt"
	"os"
)

func main() {
	mesg := []byte("This is a very secret message!")
	key := []byte("secret1")

	fmt.Printf("Message: %s\n", string(mesg))
	enc := xor(mesg, key)

	f, _ := os.Create("enc.txt")
	f.Write(enc)
	f.Close()

	// fmt.Printf("Encypted: %s\n", string(enc))

	// -----
	enc2 := xor(enc, key)
	fmt.Printf("Decrypted: %s\n", string(enc2))

}

/*
 xor performs byte-wise operation on a byte slice 'buf' using
 a second byte slice 'key'

 Assume:
   N = len(buf)
   M = len(key)

 Scenario 1: N < M
 buf:  <x0 x1 x2 x4 ... xN-1>
 key:  <k0 k1 k2 k4 ... ... |... kM-1>

 Scenario 2: N == M
 buf:  <x0 x1 x2 x4 ... xN-1>
 key:  <k0 k1 k2 k4 ... kM-1>

 Scenario 3: N > M
 buf:  <x0 x1 x2 x4 ... xM-1 xM xM+1 xM+2...>
 key:  <k0 k1 k2 k4 ... kM-1 k0 k1   k2  ...>

	<buffer here is very long>
	<key is short|key is shor>
*/
func xor(buf []byte, key []byte) []byte {
	var out []byte
	m := len(key)

	var i int
	for _, v := range buf {
		out = append(out, key[i]^v)
		i++
		if i == m {
			i = 0
		}
	}

	/*
		out := make([]byte, len(buf))
		m := len(key)

		for i, v := range buf {
			out[i] = key[i%m] ^ v
		}
	*/

	return out
}
