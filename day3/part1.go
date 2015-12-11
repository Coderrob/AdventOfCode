package main

// North 	^
// South 	v
// East 	> 
// West 	<

/*
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.
*/

import (
	"fmt"
)

func getHouseCountVistedMoreThanOnceBySanta(path string) int {
	var houseCount int
	var positions = map[string]int { "0x0": 1 }
	var xPosition int
	var yPosition int
	for index := 0; index < len(path); index++ {
		switch(path[index]){
			case '<':
				xPosition--
			case '>':
				xPosition++
			case '^':
				yPosition++
			case 'v':
				yPosition--
		}	
		
		var coordinate = fmt.Sprintf("%d,%d", xPosition, yPosition)
		
		if (positions[coordinate] != 0) {
			positions[coordinate]++
		} else {
			positions[coordinate] = 1
			houseCount++
		}
	}
	
	return houseCount
}

func getTwoHousesVisitedTestPath() string {
	return "^v^v^v^v^v"
}

func getFourHousesVisistedTestPath() string {
	return "^>v<"
}