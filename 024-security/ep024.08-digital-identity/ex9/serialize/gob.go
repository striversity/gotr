package serialize

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func GobDecode(v interface{}, buf []byte) error {
	if v == nil || buf == nil {
		return fmt.Errorf("Invalid paraemters, v nor buf cannot be nil")
	}

	b := bytes.NewReader(buf)
	dec := gob.NewDecoder(b)

	return dec.Decode(v)
}

func GobEncode(v interface{}) []byte {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(v)
	return b.Bytes()
}
