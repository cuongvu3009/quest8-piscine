package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// If no arguments provided, read from standard input
	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			for _, ch := range "ERROR: reading standard input: " {
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		for _, ch := range string(data) {
			z01.PrintRune(ch)
		}
		return
	}
	// Otherwise, read each file and print its content
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			for _, ch := range "ERROR: open " + arg + ": no such file or directory\n" {
				z01.PrintRune(ch)
			}
			os.Exit(1)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			for _, ch := range "ERROR: open " + arg + ": " + err.Error() + "\n" {
				z01.PrintRune(ch)
			}
			os.Exit(1)
		}
		for _, ch := range string(data) {
			z01.PrintRune(ch)
		}
	}
}
