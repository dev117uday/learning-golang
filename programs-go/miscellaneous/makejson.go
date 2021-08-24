package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name,address string
	fmt.Println("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Println("Enter your address")
	fmt.Scanln(&address)
	nameMap := make(map[string]string)
	nameMap = map[string]string{
		name: address,
	}
	nameJson,_ := json.Marshal(&nameMap)
	fmt.Println("Printing Json object : ",string(nameJson))
}
