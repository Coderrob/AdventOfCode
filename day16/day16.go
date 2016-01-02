package main

import "fmt"

func main() {
    runTests()
    
    properties := getTickerTapeProperties()
    auntSueList := loadAuntSueList("aunts.txt")
    
    fmt.Println("Part 1")
    auntSueID, found := findAuntSueIDByProperties(auntSueList, properties)
    fmt.Println(auntSueID, found)
    
    fmt.Println("Part 2")
    auntSueID, found = findAuntSueIDByPropertyRanges(auntSueList, properties)
    fmt.Println(auntSueID, found)
}