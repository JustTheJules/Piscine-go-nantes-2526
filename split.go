package piscine

func Split(s, sep string) []string {
	if sep == "" {
		runes := []rune(s)
		result := make([]string, len(runes))
		for i, r := range runes {
			result[i] = string(r)
		}
		return result
	}
	var result []string
	start := 0
	sepLen := len(sep)
	for {
		index := indexOf(s[start:], sep)
		if index == -1 {
			break
		}
		result = append(result, s[start:start+index])
		start += index + sepLen
	}
	result = append(result, s[start:])
	return result
}

func indexOf(s, sep string) int {
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}
