package main

import "fmt"

func main() {
	containers := loadContainerList("containers.txt")
	fmt.Println("Part 1")
	fmt.Println(getCountOfUniqueContainers(150, containers))

	fmt.Println("Part 2")
	fmt.Println(getMinimumCountOfUniqueContainers(150, containers))
}
