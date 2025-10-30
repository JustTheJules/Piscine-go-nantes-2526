package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}
	for _, arg := range args {
		num := 0
		isNumber := true
		for _, r := range arg {
			if r < '0' || r > '9' {
				isNumber = false
				break
			}
			num = num*10 + int(r-'0')
		}
		if !isNumber {
			z01.PrintRune(' ')
			continue
		}
		if num >= 1 && num <= 26 {
			if upper {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
