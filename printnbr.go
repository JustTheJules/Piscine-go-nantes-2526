package piscine

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb == -9223372036854775808 {
		for _, c := range "-9223372036854775808" {
			z01.PrintRune(c)
		}
		return
	}
	if nb < 0 {
		z01.PrintRune('-')
	} else {
		nb = -nb
	}
	if nb/10 != 0 {
		PrintNbr(-nb / 10)
	}
	digit := -(nb % 10)
	z01.PrintRune(rune(digit + '0'))
}
