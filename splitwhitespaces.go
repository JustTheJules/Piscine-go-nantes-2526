package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	start := -1
	for i, char := range s {
		if char != ' ' && char != '\t' && char != '\n' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				result = append(result, s[start:i])
				start = -1
			}
		}
	}
	if start != -1 {
		result = append(result, s[start:])
	}
	return result
}
