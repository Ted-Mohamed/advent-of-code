package main

import (
	"bufio"
	"fmt"
	"os"
)

// func isValidPasswordPartOne(input string) bool {
// 	var min int
// 	var max int
// 	var c rune
// 	var password string

// 	fmt.Sscanf(input, "%d-%d %c: %s", &min, &max, &c, &password)

// 	count := 0
// 	for _, cc := range password {
// 		if cc == c {
// 			count++
// 		}
// 	}
// 	return count >= min && count <= max
// }

func isValidPasswordPartTwo(input string) bool {
	var pos1 int
	var pos2 int
	var c rune
	var password string

	fmt.Sscanf(input, "%d-%d %c: %s", &pos1, &pos2, &c, &password)

	return password[pos1-1] != password[pos2-1] && (password[pos1-1] == byte(c) || password[pos2-1] == byte(c))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		if isValidPasswordPartTwo(scanner.Text()) {
			count++
		}
	}
	fmt.Println(count)
}
