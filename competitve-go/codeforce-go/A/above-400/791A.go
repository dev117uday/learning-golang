package main

import "fmt"

func main() {

	var limak,bob uint32 = 0,0
	_, _ = fmt.Scan(&limak)
	_, _ = fmt.Scan(&bob)

	var i int = 0

	for {
		i++
		limak = limak*3
		bob = bob*2
		if limak > bob {
			break
		}
	}
	fmt.Println(i)

}
