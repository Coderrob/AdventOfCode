package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func loadContainerList(fileName string) []int {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	containers := []int{}
	for scanner.Scan() {
		size, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, size)
	}
	return containers
}

func getPossibleContainerCombinations(target int, values []int) [][]int {
	possibleValues := pow(2, len(values))
	output := [][]int{}

	for of := possibleValues; of > 0; of-- {
		row := []int{}
		for i := 0; i < len(values); i++ {
			y := 1 << uint(i)
			if y&of == 0 && y != of {
				row = append(row, values[i])
			}
		}
		if len(row) > 0 && sum(row) == target {
			output = append(output, row)
		}
	}
	return output
}

func sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func pow(value, power int) int {
	if value == 2 {
		return 1 << uint(power)
	}

	return int(math.Pow(float64(value), float64(power)))
}

func getMinimumLengthOfContainers(combinations [][]int) int {
	minLength := 1 << 32
	for _, containers := range combinations {
		if len(containers) < minLength {
			minLength = len(containers)
		}
	}
	return minLength
}

func getCountOfContainersByLength(combinations [][]int, length int) int {
	count := 0
	for _, containers := range combinations {
		if len(containers) == length {
			count++
		}
	}
	return count
}
