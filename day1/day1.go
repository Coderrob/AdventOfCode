package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 1")
	fmt.Println(getSantasFloor(getSantasPath()))
	
	fmt.Println("Part 2")
	fmt.Println(getPositionWhenSantasGoesIntoTheBasement(getSantasPath()))
}