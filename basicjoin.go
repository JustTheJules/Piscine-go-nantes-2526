package piscine

func BasicJoin(elems []string) string {
	result := ""
	for i := 0; i < len(elems); i++ {
		result += elems[i]
	}
	return result
}
