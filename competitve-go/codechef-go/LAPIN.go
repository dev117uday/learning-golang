package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	var numberOfIntegers uint8
	var vector []string
	_, _ = fmt.Scanln(&numberOfIntegers)

	// input strings into a vector
	{
		var i uint8 = 0
		var number string
		for i = 0 ; i< numberOfIntegers; i++ {
			_, _ = fmt.Scanln(&number)
			vector = append(vector,number)
		}
	}


	for _,value := range vector {
		length := len(value)
		if length%2==0{

			first := value[length/2:]
			temp := strings.Split(first,"")
			sort.Strings(temp)
			first = strings.Join(temp,"")

			second := value[:length/2]
			temp = strings.Split(second,"")
			sort.Strings(temp)
			second = strings.Join(temp,"")

			if first == second {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}

		} else {

			first := value[:length/2]
			temp := strings.Split(first,"")
			sort.Strings(temp)
			first = strings.Join(temp,"")

			second := value[length/2+1:]
			temp = strings.Split(second,"")
			sort.Strings(temp)
			second = strings.Join(temp,"")

			if first == second {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}


}
