package main

import "fmt"

func main() {
    runTests()
    return
    
    containers := loadContainerList("containers.txt")
    fmt.Println("Part 1")
    fmt.Println(getCountOfContainerCombinationsByLiter(150, containers))
}