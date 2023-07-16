package main

import (
	"encoding/base64"
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
	payload := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Payload: %s\n", payload)
}
