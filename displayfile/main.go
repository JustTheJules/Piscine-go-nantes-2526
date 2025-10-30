package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(string(content))
}
