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

	for scanner.Scan() {
		line := scanner.Text()
		offsetX += 3
		if line[offsetX%len(line)] == byte('#') {
			count++
		}
	}

	println(count)
}
