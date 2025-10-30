package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return nb
	} else {
		for i := 1; i <= nb/2; i++ {
			if i*i == nb {
				return i
			}
		}
	}
	return 0
}
