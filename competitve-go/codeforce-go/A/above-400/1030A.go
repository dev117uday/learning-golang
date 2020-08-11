package main

import "fmt"

func main() {

	var number int
	_, _ = fmt.Scanln(&number)
	var oned = make([]int,number)
	flag := 0
	for i:=0;i<len(oned); i++ {
		_, _ = fmt.Scan(&oned[i])
		if oned[i] == 1 {
			flag = 1
		}
	}
	 if flag == 0 {
	 	fmt.Print("EASY")
	 } else {
	 	fmt.Print("HARD")
	 }


}
