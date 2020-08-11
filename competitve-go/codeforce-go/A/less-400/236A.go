package main

import "fmt"

func main() {

	var inputString string
	_, _ = fmt.Scanln(&inputString)

	var hashMap  = make(map[string]string)

	temp := string(inputString[0])

	hashMap[temp] = string(inputString[0])

	for i:=1; i<len(inputString); i++ {
		_,found := hashMap[string(inputString[i])]
		if found {
			continue
		} else {
			hashMap[string(inputString[i])] = string(inputString[i])
		}
	}
	if len(hashMap)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}


}
