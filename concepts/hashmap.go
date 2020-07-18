package main

import (
	"fmt"
)

func main()  {

	var idMap map[string]int = make(map[string]int)
	idMap = map[string]int {	"joe":123	}

	fmt.Println("idMap : ",idMap)
	fmt.Println("value :", idMap["joe"] )
	delete(idMap,"joe")
	_id, p := idMap["joe"]
	fmt.Println("id : ",_id," p : ",p)
	fmt.Println(len(idMap))
	idMap = map[string]int {	"uday":999	}
	for key,val := range idMap {
		fmt.Println(key,val)
	}
}