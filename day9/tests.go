package main

import (
	"os"
	"bufio"
)

func getTestRoutesFromFile() []Route {
	file, _ := os.Open("testpaths.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var routes = []Route {} 
	for scanner.Scan() {
		routes = append(routes, getRouteFromString(scanner.Text()))
	} 
	return routes
}