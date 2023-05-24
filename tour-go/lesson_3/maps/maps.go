package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Delhi": {3.22, 5.22},
	"Goa":   {3.9034, 23.232},
}

func main() {

	var m map[string]int8
	m = make(map[string]int8)
	// or use :=

	m["Uday"] = 19
	m["Goa"] = 1
	fmt.Println("Uday : ", m["Uday"])

	x := make(map[string]Vertex)
	x["Delhi"] = Vertex{
		3.22, 4.21,
	}
	fmt.Println(x["Delhi"])

	fmt.Println(m)
	delete(m, "Goa")
	fmt.Println(m)

}
