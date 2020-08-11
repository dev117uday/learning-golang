package main

import "fmt"

func main(){

	var numberOfRows int = 0
	_,_ = fmt.Scanln(&numberOfRows)

	var twoDMatrix = make([][3]int,numberOfRows)

	for i:=0; i<numberOfRows; i++ {
		for j:=0; j<3; j++ {
			_,_ = fmt.Scan(&twoDMatrix[i][j])
		}
	}

	var count int = 0
	for i:=0; i<numberOfRows; i++ {
		temp := 0
		for j:=0; j<3; j++ {
			if twoDMatrix[i][j] == 1 {
				temp++
			}
		}
		if temp > 1 {
			count++
		}
	}
	
	fmt.Println(count)

}