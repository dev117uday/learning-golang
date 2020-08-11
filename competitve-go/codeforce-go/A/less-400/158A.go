package main

import (
	"fmt"
	"os"
)

func main() {

	var numberOfParticipants, rating int = 0, 0
	_, _ = fmt.Scan(&numberOfParticipants)
	_, _ = fmt.Scan(&rating)

	var slice = make([]int, numberOfParticipants)
	for i := 0; i < numberOfParticipants; i++ {
		_, _ = fmt.Scan(&slice[i])
	}

	if slice[0] == 0 {
		fmt.Println(0)
		os.Exit(0)
	}

	score := slice[rating-1]
	count := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] >= score && slice[i] != 0 {
			count++
		} else {
			break
		}
	}

	fmt.Println(count)

}
