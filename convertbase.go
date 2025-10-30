package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if !isValidBaseBis(baseFrom) || !isValidBaseBis(baseTo) {
		return ""
	}
	decimalValue := toDecimal(nbr, baseFrom)
	if decimalValue == -1 {
		return ""
	}
	return fromDecimal(decimalValue, baseTo)
}

func isValidBaseBis(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || char <= 32 || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func toDecimal(nbr, base string) int {
	baseLen := len(base)
	charIndex := make(map[rune]int)
	for i, char := range base {
		charIndex[char] = i
	}
	result := 0
	for _, char := range nbr {
		index, exists := charIndex[char]
		if !exists {
			return -1
		}
		result = result*baseLen + index
	}
	return result
}

func fromDecimal(num int, base string) string {
	if num == 0 {
		return string(base[0])
	}
	baseLen := len(base)
	var result []rune
	for num > 0 {
		remainder := num % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		num /= baseLen
	}
	return string(result)
}
