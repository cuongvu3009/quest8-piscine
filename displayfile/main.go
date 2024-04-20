package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	// Check if the filename is missing
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	// Check if too many arguments are provided
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	// Read the content of the file
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	// Print the content of the file to standard output
	fmt.Print(string(content))
}
