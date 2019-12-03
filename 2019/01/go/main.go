package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func inputFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var values []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			values = append(values, value)
		}
	}
	return values, scanner.Err()
}

func fuelForMass(m int) int {
	return int(math.Floor(float64(m)/3.0)) - 2
}

func main() {
	input, err := inputFromFile("2019/01/go/input")

	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, v := range input {
		sum += fuelForMass(v)
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	for _, v := range input {
		mass := fuelForMass(v)
		for mass > 0 {
			sum += mass
			mass = int(math.Floor(float64(mass)/3.0)) - 2
		}
	}
	fmt.Println("Part 2:", sum)

}
