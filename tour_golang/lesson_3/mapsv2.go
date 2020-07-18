package main

import "fmt"

type Vertex struct {
	Lat,Long float64
}

func main(){
	var m map[string]int8
	m = make(map[string]int8)
	m["Uday"] = 19
	fmt.Println("Uday : ",m["Uday"])
	var x map[string]Vertex
	x = make(map[string]Vertex)
	x["Delhi"] = Vertex{
		3.22,4.21,
	}
	fmt.Println(x["Delhi"])

}