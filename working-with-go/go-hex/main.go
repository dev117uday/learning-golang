package main

import (
	"fmt"
	"go-hex/pkg/controller"
	"go-hex/pkg/service"
	"go-hex/pkg/storage"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	r, err := storage.SetupStorage()
	if err != nil {
		fmt.Println(err)
	}

	s := service.NewService(r)

	route := controller.Handler(s)

	log.Fatal(http.ListenAndServe(":8080", route))

}
