package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for i, r := range arguments {
		if len(os.Args) >= 1 {
			if i > 1 {
				z01.PrintRune(r)
			}
		}
	}
	z01.PrintRune('\n')
}
