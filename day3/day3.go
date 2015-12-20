package main

import (
	"fmt"
)

func main() {
    path := getSantasPath()
    
	fmt.Println("Part 1")	
	fmt.Println(getCountOfHousesVisitedMoreThanOnce(path))
	
	fmt.Println("Part 2")
	fmt.Println(getCountOfHousesVisitedMoreThanOnceBySantaAndRobot(path))
}