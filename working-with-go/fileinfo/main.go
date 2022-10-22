package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var path string
var dashes string = "\n-----------------------------------"

func main() {

	fmt.Print(dashes)

	commandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	commandLine.StringVar(&path, "p", "", "help message for flagname")

	err := commandLine.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error in parsing command line arguments")
		return
	}

	if path == "" {
		fmt.Print("\nplease provide argument to -p")
		fmt.Print(dashes)
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("file/directory not found")
		fmt.Print(dashes)
		return
	}

	mode := info.Mode()

	unixTimeUTC := time.Unix(1405544146, 0)
	fmt.Printf("\nLastModified : %s\n", unixTimeUTC)

	if info.IsDir() {
		fmt.Printf("DirectoryName : %s\n", info.Name())
		fmt.Println("Its is a directory")
	} else {
		fmt.Printf("FileName : %s\n", info.Name())
		size := fmt.Sprintf("%d", info.Size())
		mb := float64(info.Size()) / 1000 / 1000
		fmt.Printf("File Size (bytes): %s\n", size)
		fmt.Printf("File Size (mb)   : %.2f\n", mb)
	}

	fmt.Print("Owner: ")
	for i := 1; i < 4; i++ {
		fmt.Print(string(mode.String()[i]))
	}

	fmt.Print("\nGroup: ")
	for i := 4; i < 7; i++ {
		fmt.Print(string(mode.String()[i]))
	}

	fmt.Print("\nOther: ")
	for i := 7; i < 10; i++ {
		fmt.Print(string(mode.String()[i]))
	}

	fmt.Print(dashes)
}
