package piscine

func ToLower(s string) string {
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			s = s[:i] + string(char+32) + s[i+1:]
		}
	}
	return s
}
