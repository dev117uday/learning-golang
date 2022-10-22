package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwt_key = []byte("this is my jwt secret key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/home", home)

	fmt.Println("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.UserName]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader((http.StatusUnauthorized))
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		UserName: credentials.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwt_key)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, tokenString)

}

func mainx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "mainx")
}

func home(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")[7:]

	fmt.Println(tokenString)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	fmt.Printf("%+v", tkn)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(tkn.Valid)

	if !tkn.Valid {
		w.WriteHeader((http.StatusUnauthorized))
		return
	}

	fmt.Fprintf(w, "hello, %s", claims.UserName)

}

func refresh(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")[7:]

	fmt.Println(tokenString)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwt_key)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
