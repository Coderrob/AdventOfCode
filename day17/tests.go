package main

import "fmt"

func runTests() {
	containers := loadContainerList("examples.txt")
	fmt.Println(getCountOfUniqueContainers(25, containers))
}
