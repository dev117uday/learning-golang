package main

import "fmt"

func main() {

	var elements = make([]map[string]string,1)
	elements[0] = map[string]string {
		"uday" : "uday",
	}

	fmt.Println(elements[0]["uday"])


}
