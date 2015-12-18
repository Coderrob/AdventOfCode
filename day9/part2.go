package main

/*
--- Part Two ---

The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?
*/

func getLongestPathDistance(routes []Route) int {
	cityNames := getUniqueCityNames(routes)
	var longestDistance int
	
	routePermutation, _ := NewPerm(cityNames, nil)
	
	for permutation, err := routePermutation.Next(); err == nil; permutation, err = routePermutation.Next() {
		distance := getPathDistance(permutation.([]string), routes)
		if longestDistance == 0 || longestDistance < distance {
			longestDistance = distance
		}
	}
	
	return longestDistance
}