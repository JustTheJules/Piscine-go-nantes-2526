package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		for _, char := range str {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
