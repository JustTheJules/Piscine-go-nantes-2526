package main

import "github.com/01-edu/z01"

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func printPoint(x, y int) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(y)
	z01.PrintRune('\n')
}

func main() {
	var x, y int
	setPoint(&x, &y)
	printPoint(x, y)
}
