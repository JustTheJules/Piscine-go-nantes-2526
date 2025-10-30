package piscine

func NRune(s string, n int) rune {
	if n-1 < 0 || n-1 >= len([]rune(s)) {
		return 0
	}
	return []rune(s)[n-1]
}
