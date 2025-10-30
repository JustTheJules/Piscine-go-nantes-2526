package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	mainStr := ""

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg
		}
	}

	result := mainStr + insertStr
	if order {
		result = sortString(result)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
