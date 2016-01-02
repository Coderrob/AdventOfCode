package main

import "fmt"

func runTests() {
    containers := loadContainerList("examples.txt")
    fmt.Println(getCountOfContainerCombinationsByLiter(20, containers))
}