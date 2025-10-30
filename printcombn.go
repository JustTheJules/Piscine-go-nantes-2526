package piscine

import "github.com/01-edu/z01"

func PrintCombN(nb int) {
	if nb < 1 || nb > 9 {
		z01.PrintRune('\n')
		return
	}
	digits := make([]rune, nb)
	for i := 0; i < nb; i++ {
		digits[i] = rune('0' + i)
	}

	for {
		for _, d := range digits {
			z01.PrintRune(d)
		}
		if digits[0] == rune('9'-nb+1) && digits[nb-1] == '9' {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')

		for i := nb - 1; i >= 0; i-- {
			if digits[i] < rune('9'-(nb-1-i)) {
				digits[i]++
				for j := i + 1; j < nb; j++ {
					digits[j] = digits[j-1] + 1
				}
				break
			}
		}

	}
	z01.PrintRune('\n')
}
