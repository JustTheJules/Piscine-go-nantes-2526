package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i, arg := range arguments {
		if i > 0 {
			for _, r := range arg {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
		}
	}
	if len(arguments) == 1 {
		z01.PrintRune('\n')
	}
}
