package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	input, err := inputToIntArray(os.Stdin)

	for i, v1 := range input {
		for j, v2 := range input {
			for k, v3 := range input {
				if j != i && k != j && k != i && v1+v2+v3 == 2020 {
					fmt.Println(v1 * v2 * v3)
					return
				}
			}
		}
	}

	if err != nil {
		fmt.Println(err)
	}

}
