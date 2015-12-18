package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Route struct {
	current string
	destination string
	distance int
}

func getRoutesFromFile() []Route {
	file, _ := os.Open("paths.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var routes = []Route {} 
	for scanner.Scan() {
		routes = append(routes, getRouteFromString(scanner.Text()))
	} 
	return routes
}

func getRouteFromString(input string) Route {		
	routeParts := strings.Split(input, " to ")	
	destinationParts := strings.Split(routeParts[1], " = ")
	
	startingCity := routeParts[0]
	destinationCity := destinationParts[0]
	distance, _ := strconv.Atoi(destinationParts[1])
	
	var route = Route {}
	route.current = startingCity
	route.destination = destinationCity
	route.distance = distance
	return route
}

func getPathDistance(route []string, routes []Route) int {
	var distance int
	for index := 1; index < len(route); index++ {		
		var startingCity = route[index - 1]
		var desintation = route[index]		
		for _, route := range routes {
			if (route.current == startingCity && route.destination == desintation) || 
				route.current == desintation && route.destination == startingCity {
				distance += route.distance
			}
		}
	}
	return distance
}

func getUniqueCityNames(routes []Route) []string {
	var cityNamesMap = map[string]bool {}
	for _, route := range routes {
		if cityNamesMap[route.current] == false {
			cityNamesMap[route.current] = true
		}
		if cityNamesMap[route.destination] == false {
			cityNamesMap[route.destination] = true
		}
	}	
	return getKeysFromStringMap(cityNamesMap)
}

func getKeysFromStringMap(input map[string]bool) []string {
	keys := make([]string, 0, len(input))
    for k := range input {
        keys = append(keys, k)
    }
	return keys
}