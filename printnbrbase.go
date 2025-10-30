package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	baseLen := len(base)
	isNegative := nbr < 0
	if isNegative {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			printPositiveNbr(uint64(nbr), base, baseLen)
			return
		}
		nbr = -nbr
	}
	printPositiveNbr(uint64(nbr), base, baseLen)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '-' || c == '+' || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func printPositiveNbr(nbr uint64, base string, baseLen int) {
	if nbr == 0 {
		return
	}
	printPositiveNbr(nbr/uint64(baseLen), base, baseLen)
	z01.PrintRune(rune(base[nbr%uint64(baseLen)]))
}
