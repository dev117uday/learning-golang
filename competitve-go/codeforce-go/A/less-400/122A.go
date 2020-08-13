package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	flag := 0
	var check = []int{4,7,47,74,44,444,447,474,477,777,774,744}
	for i:=0; i<len(check); i++ {
		if num%check[i] == 0 {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
