package main

import "fmt"

func main() {

	var cost,number, amount int32
	_, _ = fmt.Scan(&cost)
	_, _ = fmt.Scan(&amount)
	_, _ = fmt.Scan(&number)

	temp := cost*(number*(number+1))/2

	if temp - amount < 0 {
		fmt.Println("0")
	} else {
		final := temp-amount
		fmt.Println(final)
	}

}
