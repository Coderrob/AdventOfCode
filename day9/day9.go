package main

import "fmt"

func main() {
	routes := getRoutesFromFile()
	
	fmt.Println("Part 1")
	shortestDistance := getShortedPathDistance(routes)
	fmt.Println("shortest path", shortestDistance)
	
	fmt.Println("Part 2")
	longestDistance := getLongestPathDistance(routes)
	fmt.Println("longest path", longestDistance)
}