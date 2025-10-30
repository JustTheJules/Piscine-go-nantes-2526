package piscine

func TrimAtoi(s string) int {
	var positif bool = true
	var numberdetected bool = false
	var i int

	for _, c := range s {
		if c == '-' && !numberdetected {
			positif = false
		}
		if c <= '9' && c >= '0' {
			numberdetected = true
			i = i*10 + int(c-'0')
		}
	}

	if positif {
		return i
	} else {
		return -i
	}
}
