package main

import "fmt"

func runTests() {
    auntSueList := loadAuntSueList("aunts.txt")
    fmt.Println(auntSueList[0])
}