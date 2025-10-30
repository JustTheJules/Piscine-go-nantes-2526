package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1:]
	vowels := "aeiouAEIOU"
	var extractedVowels []rune

	for _, word := range input {
		for _, char := range word {
			if contains(vowels, char) {
				extractedVowels = append([]rune{char}, extractedVowels...)
			}
		}
	}

	vowelIndex := 0
	for i, word := range input {
		for _, char := range word {
			if contains(vowels, char) {
				z01.PrintRune(extractedVowels[vowelIndex])
				vowelIndex++
			} else {
				z01.PrintRune(char)
			}
		}
		if i < len(input)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func contains(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
