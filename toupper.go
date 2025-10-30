package piscine

func ToUpper(s string) string {
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			s = s[:i] + string(char-32) + s[i+1:]
		}
	}
	return s
}
