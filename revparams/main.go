package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i >= 1; i-- {
		for _, r := range arguments[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
	if len(arguments) == 1 {
		z01.PrintRune('\n')
	}
}
