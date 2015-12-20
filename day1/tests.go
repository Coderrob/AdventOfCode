package main

import "fmt"

func runDayOneTests() {
    fmt.Println(getSantasFloor("(())") == 0)
    fmt.Println(getSantasFloor(")())())") == -3)
    
    fmt.Println(getPositionWhenSantasGoesIntoTheBasement(")") == 1)
    fmt.Println(getPositionWhenSantasGoesIntoTheBasement("()())") == 5)
}