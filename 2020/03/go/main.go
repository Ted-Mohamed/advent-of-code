package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	offsetX := 0
	scanner.Scan()

	incX := 1

	for scanner.Scan() {
		// scanner.Scan() // incY will be scaning n-1 extra times before reading line
		line := scanner.Text()
		offsetX += incX
		if line[offsetX%len(line)] == byte('#') {
			count++
		}
	}

	println(count)
}
