package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	digits := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return 0
		}
		if _, exists := digits[base[i]]; exists {
			return 0
		}
		digits[base[i]] = i
	}
	result := 0
	for i := 0; i < len(s); i++ {
		digit, exists := digits[s[i]]
		if !exists {
			return 0
		}
		result = result*len(base) + digit
	}
	return result
}
