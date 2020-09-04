package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	sender()
	reciever()
	attacker()
}
func sender() {
	mesg := []byte("This is a very secret message!")

	fmt.Printf("Message: %s\n", string(mesg))
	enc := encode(mesg)

	f, _ := os.Create("enc.txt")
	f.Write(enc)
	f.Close()
}

func reciever() {
	enc, _ := ioutil.ReadFile("enc.txt")

	enc2 := decode(enc)
	fmt.Printf("Decoded: %s\n", string(enc2))
}

func attacker() {
	enc, _ := ioutil.ReadFile("enc.txt")

	enc2 := decode(enc)
	fmt.Printf("Attacker: %s\n", string(enc2))
}

func encode(buf []byte) []byte {
	return bytes.Map(func(in rune) rune {
		return em[in]
	}, buf)
}
func decode(buf []byte) []byte {
	return bytes.Map(func(in rune) rune {
		return dm[in]
	}, buf)
}

var em = map[rune]rune{
	'q': 'a', 'w': 'b', 'e': 'c', 'r': 'd', 't': 'e', 'y': 'f',
	'u': 'g', 'i': 'h', 'o': 'i', 'p': 'j', 'a': 'k', 's': 'l',
	'd': 'm', 'f': 'n', 'g': 'o', 'h': 'p', 'j': 'q', 'k': 'r',
	'l': 's', 'z': 't', 'x': 'u', 'c': 'v', 'v': 'w', 'b': 'x',
	'n': 'y', 'm': 'z',
	'Q': 'A', 'W': 'B', 'E': 'C', 'R': 'D', 'T': 'E', 'Y': 'F',
	'U': 'G', 'I': 'H', 'O': 'I', 'P': 'J', 'A': 'K', 'S': 'L',
	'D': 'M', 'F': 'N', 'G': 'O', 'H': 'P', 'J': 'Q', 'K': 'R',
	'L': 'S', 'Z': 'T', 'X': 'U', 'C': 'V', 'V': 'W', 'B': 'X',
	'N': 'Y', 'M': 'Z',
	' ': '_',
}

var dm = map[rune]rune{
	'a': 'q', 'b': 'w', 'c': 'e', 'd': 'r', 'e': 't', 'f': 'y',
	'g': 'u', 'h': 'i', 'i': 'o', 'j': 'p', 'k': 'a', 'l': 's',
	'm': 'd', 'n': 'f', 'o': 'g', 'p': 'h', 'q': 'j', 'r': 'k',
	's': 'l', 't': 'z', 'u': 'x', 'v': 'c', 'w': 'v', 'x': 'b',
	'y': 'n', 'z': 'm',
	'A': 'Q', 'B': 'W', 'C': 'E', 'D': 'R', 'E': 'T', 'F': 'Y',
	'G': 'U', 'H': 'I', 'I': 'O', 'J': 'P', 'K': 'A', 'L': 'S',
	'M': 'D', 'N': 'F', 'O': 'G', 'P': 'H', 'Q': 'J', 'R': 'K',
	'S': 'L', 'T': 'Z', 'U': 'X', 'V': 'C', 'W': 'V', 'X': 'B',
	'Y': 'N', 'Z': 'M',
	'_': ' ',
}
