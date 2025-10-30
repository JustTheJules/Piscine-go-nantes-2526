package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printError(err.Error())
			os.Exit(1)
		}

		io.Copy(os.Stdout, file)
		file.Close()
	}
}

func printError(message string) {
	z01.PrintRune('E')
	z01.PrintRune('R')
	z01.PrintRune('R')
	z01.PrintRune('O')
	z01.PrintRune('R')
	z01.PrintRune(':')
	z01.PrintRune(' ')
	for _, char := range message {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
