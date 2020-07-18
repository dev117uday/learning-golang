package main

import "fmt"

type Vertex struct {
	Lat,Long float64
}

var m = map[string]Vertex {
	"Delhi" : {	3.22,5.22,	},
	"Goa" : {	3.9034,23.232,	},
}

func main()  {
	fmt.Println(m)
	delete(m,"Goa")
	fmt.Println(m)
}