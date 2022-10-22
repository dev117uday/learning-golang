package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("myprivatekey")

func main() {
	fmt.Println("hello World")

	jwtToken, err := GenerateJwt()

	if err != nil {
		fmt.Printf("err %s", err)
	}

	fmt.Println(isAuthorized(jwtToken))

}

func GenerateJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = "Uday Yadav"
	claims["exp"] = time.Now().Add(time.Second * 1)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", fmt.Errorf("something went wrong %s", err)
	}

	return tokenString, nil

}

func isAuthorized(token string) bool {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}

	claims := tkn.Claims.(jwt.MapClaims)

	s := claims["exp"].(string)

	fmt.Println(s)

	t, _ := time.Parse("RFC3339", s)

	fmt.Println(t)

	if time.Now().Add(time.Second * 10).After(t) {
		return true
	} else {
		return false
	}

}
