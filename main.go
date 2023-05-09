package main

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:embed key.pem
var s []byte

//go:embed public.pem
var ss []byte

func main() {

	private, err := jwt.ParseRSAPrivateKeyFromPEM(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	public, err := jwt.ParseRSAPublicKeyFromPEM(ss)
	if err != nil {
		log.Fatal(err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"foo": "bar",
			"exp": time.Now().Add(8 * time.Hour).Unix(),
		},
	)

	token2, err := token.SignedString(private)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("token" + token2)

	result, err := jwt.ParseWithClaims(token2, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return public, nil
	})

	fmt.Println(result.Claims)
	if err != nil {
		log.Fatal(err.Error())
	}
}
