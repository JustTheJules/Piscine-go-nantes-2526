package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	lenStr := len(str)
	for i := 0; i < lenStr; i++ {
		if !((str[i] <= 'z' && str[i] >= 'a') || (str[i] <= 'Z' && str[i] >= 'A') || (str[i] <= '9' && str[i] >= '0')) {
			return false
		}
	}
	return true
}
