package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Fprintln(os.Stderr, "Usage: go run . -c <number> <file1> [<file2> ...]")
		os.Exit(1)
	}

	count, err := parseCount(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	files := os.Args[3:]
	exitCode := 0
	printed := false

	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
			printed = true
			continue
		}

		info, err := f.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
			printed = true
			f.Close()
			continue
		}

		if printed {
			fmt.Println()
		}

		fmt.Printf("==> %s <==\n", fname)

		size := info.Size()
		start := size - int64(count)
		if start < 0 {
			start = 0
		}

		buf := make([]byte, size-start)
		_, err = f.ReadAt(buf, start)
		f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitCode = 1
			printed = true
			continue
		}

		fmt.Print(string(buf))
		printed = true
	}

	os.Exit(exitCode)
}

func parseCount(arg string) (int, error) {
	count := 0
	for _, ch := range arg {
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("invalid count: %s", arg)
		}
		count = count*10 + int(ch-'0')
	}
	return count, nil
}
