package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	length := 0
	temp := n
	for temp > 0 {
		length++
		temp /= 10
	}
	tab := make([]int, length)
	for nbr := length - 1; nbr >= 0; nbr-- {
		tab[nbr] = n % 10
		n /= 10
	}
	for i := 1; i < len(tab); i++ {
		key := tab[i]
		j := i - 1
		for j >= 0 && tab[j] > key {
			tab[j+1] = tab[j]
			j = j - 1
		}
		tab[j+1] = key
	}
	for k := 0; k < len(tab); k++ {
		PrintNbr(tab[k])
	}
}
