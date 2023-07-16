package main

import (
	"flag"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Id    int      `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Email string   `json:"email,omitempty"`
	Iat   int64    `json:"iat,omitempty"`
	jwt.RegisteredClaims
}

var (
	key = "very-secure"
)

func init() {
	flag.StringVar(&key, "key", key, "key to sign JWT")
	flag.Parse()
}

func main() {
	claims := &MyClaims{
		Id:    1104,
		Name:  "Jane Doe",
		Roles: []string{"admin", "user"},
		Email: "jane.doe@example.com",
		Iat:   1577836800,
	}

	jwtString, err := createSignedJWT(key, claims)
	if err != nil {
		log.Fatal("unable to create signed JWT", "error", err)
	}

	fmt.Println(jwtString)

}

func createSignedJWT(key string, claims *MyClaims) (string, error) {
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token using a key
	jwtString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
