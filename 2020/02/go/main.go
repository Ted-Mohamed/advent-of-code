package main

import (
	"bufio"
	"io"
	"strconv"
)

func inputToIntArray(r io.Reader) ([]int, error) {
	var values []int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			values = append(values, value)
		}
	}
	return values, scanner.Err()
}

func main() {
}
