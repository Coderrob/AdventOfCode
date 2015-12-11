package main

import (
	"strings"
	"strconv"
)

func convertStringArrayToIntArray(strs []string) []int {
    var ints = []int{}
    for _, i := range strs {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        ints = append(ints, j)
    }
	return ints
}

func getDimentionsFromString(str string) []int {
	var values = strings.Split(str, "x")
	return convertStringArrayToIntArray(values)
}

func getSmallestSize(sizes []int) int {
	min := sizes[0]
	for _, value := range sizes {
		if value < min {
			min = value
		}
	}
	return min
}

func getTwoSmallestSizes(sizes []int) []int {
	var min1 = sizes[0]
	var min2 = sizes[1]
	
	if (min2 < min1) {
		min1 = sizes[1];
		min2 = sizes[0];
	}

	for index := 2; index < len(sizes); index++ {
		if (sizes[index] < min1) {
			min2 = min1
			min1 = sizes[index]
		} else if (sizes[index] < min2) {
			min2 = sizes[index]
		}
	}
	
	return []int { min1, min2 }
}