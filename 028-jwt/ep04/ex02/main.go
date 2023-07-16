package main

import (
	"flag"
	"fmt"
	"os"

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
	key       = "very-secure"
	jwtString string
)

func init() {
	flag.StringVar(&key, "key", key, "key to sign JWT")
	flag.Parse()
	jwtString, _ = os.LookupEnv("JWT")
}

func main() {
	log.Info("JWT String", "value", jwtString)

	claims, err := getJWTClaims(key, jwtString)
	if err != nil {
		log.Fatal("unable to get JWT claims", "error", err)
	}

	fmt.Println("Id:", claims.Id)
	fmt.Println("Name:", claims.Name)
	fmt.Println("Roles:", claims.Roles)
	fmt.Println("Email:", claims.Email)
	fmt.Println("Issued At:", claims.Iat)
}

func getJWTClaims(key string, jwtString string) (*MyClaims, error) {
	// jwtKeyFunc closure provides the key for the JWT parser signature verification
	jwtKeyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}

	// parse the JWT string and extarct the claims
	token, err := jwt.ParseWithClaims(jwtString, &MyClaims{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	// verify the token contains the expected MyClaims structure
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type")
	}

	return claims, nil
}
