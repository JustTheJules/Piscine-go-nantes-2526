package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	sLen := len(s)
	toFindLen := len(toFind)
	for i := 0; i <= sLen-toFindLen; i++ {
		j := 0
		for j < toFindLen {
			if s[i+j] != toFind[j] {
				break
			}
			j++
		}
		if j == toFindLen {
			return i
		}
	}
	return -1
}
