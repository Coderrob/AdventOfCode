package main

import "fmt"

func runDayThreeTests() {
    fmt.Println(getCountOfHousesVisitedMoreThanOnce("^v^v^v^v^v") == 2)
    fmt.Println(getCountOfHousesVisitedMoreThanOnce("^>v<") == 4)
    
    fmt.Println(getCountOfHousesVisitedMoreThanOnceBySantaAndRobot("^v") == 3)
    fmt.Println(getCountOfHousesVisitedMoreThanOnceBySantaAndRobot("^>v<") == 3)
    fmt.Println(getCountOfHousesVisitedMoreThanOnceBySantaAndRobot("^v^v^v^v^v") == 11)
}