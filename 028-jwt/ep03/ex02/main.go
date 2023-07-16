package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var data = `
	{
		"id": 1104,
		"name": "Jane Doe",
		"roles": ["admin", "user"],
		"email": "jane.doe@example.com",
		"iat": 1577836800
	}
`

func main() {
	var buf = new(bytes.Buffer)
	json.Compact(buf, []byte(data))
	payload := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Data(compact): %s\n", buf.Bytes())
	fmt.Printf("Payload: %s\n", payload)
}
