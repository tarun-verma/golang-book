package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) // I think you just pass a file descriptor here
	for input.Scan() {                  // this loop will break after passing Ctrl+D (EOF) since NewScanner reads till EOF is encountered.
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}
