package main

import "fmt"

func updatePosition2(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
		}
	}()

	return positionChannel
}

func main() {
	positionChannel1 := updatePosition2("Legolas :")
	positionChannel2 := updatePosition2("Gandalf :")

	for i := 0; i < 5; i++ {
		fmt.Println(<-positionChannel1)
		fmt.Println(<-positionChannel2)
	}

	fmt.Println("Done with getting updates on positions.")
}
