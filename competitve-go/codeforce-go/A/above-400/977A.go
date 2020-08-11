package main

import "fmt"

func main() {

	var number uint32 = 0
	var k int = 0
	_, _ = fmt.Scan(&number)
	_, _ = fmt.Scan(&k)

	for i:=0; i<k; i++ {
		if number%10 !=0 {
			number--
		} else {
			number = number/10
		}
	}

	fmt.Println(number)


}
