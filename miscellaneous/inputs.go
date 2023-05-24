package main

import (
	"bufio"
	"fmt"
	"os"
)

func intInput() {
	var number int32
	_, _ = fmt.Scanf("%d", &number)
	fmt.Println(number)
}

func StringInput() {
	fmt.Print("Enter your full name : ")
	myString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("myString : ", myString)
}

func vectorInput() {
	numberOfStrings := 0
	var vector []string
	_, _ = fmt.Scanln(&numberOfStrings)

	{
		var number string
		for i := 0; i < numberOfStrings; i++ {
			_, _ = fmt.Scanln(&number)
			vector = append(vector, number)
		}
	}
}

func twoDvectorinput() {

	var cols, rows = 0, 0
	fmt.Print("Enter the number of cols : ")
	_, _ = fmt.Scan(&cols)
	fmt.Print("Enter the number of rows : ")
	_, _ = fmt.Scan(&rows)

	var twodslices = make([][]int, rows)
	var i int
	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&twodslices[i][j])
		}
	}
	fmt.Println(twodslices)
}
