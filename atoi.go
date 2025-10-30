package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	result := 0
	for i, char := range s {
		if i == 0 {
			if char == '-' {
				sign = -1
				continue
			} else if char == '+' {
				continue
			}
		}
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result * sign
}
