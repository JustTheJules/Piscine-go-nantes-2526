package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err1 := strconv.Atoi(os.Args[1])
	b, err2 := strconv.Atoi(os.Args[3])
	op := os.Args[2]

	if err1 != nil || err2 != nil || len(op) != 1 {
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Stop : division by zero")
		} else {
			fmt.Println(a / b)
		}
	case "%":
		if b == 0 {
			fmt.Println("Stop : modulo by zero")
		} else {
			fmt.Println(a % b)
		}
	default:
		return
	}
}
