package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scanln(&num)
	var twodslices = make([][3]int, num)
	for i:=0; i<num; i++ {
		for j:=0; j<3; j++ {
			_, _ = fmt.Scan(&twodslices[i][j])
		}
	}
	flag := 0
	for i:=0; i<3; i++ {
		count :=0
		for j:=0; j<num; j++ {
			count = count + twodslices[j][i]
		}
		if count != 0 {
			fmt.Println("NO")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("YES")
	}
}
