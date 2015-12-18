package main

import (
	"strings"
)

func getNumberOfCharacters(input string) int {
	var characterCount int	
	var stringConents = input
	stringConents = strings.TrimPrefix(stringConents, "\"")		
	stringConents = strings.TrimSuffix(stringConents, "\"")
	stringConents = strings.Replace(stringConents, " ", "", -1)
	
	// Count and remove the \" from the string
	stringConents, quoteCount := getCountOfAndRemoveSpecifiedSubString(stringConents, "\\\"")
	characterCount += quoteCount
	
	stringConents, slashCount := getCountOfAndRemoveSpecifiedSubString(stringConents, "\\\\")	
	characterCount += slashCount
	
	stringConents, hexCount := getCountOfAndRemoveHexCharacters(stringConents)
	characterCount += hexCount
	
	return characterCount + len(stringConents)
}

func getCountOfAndRemoveHexCharacters(input string) (string, int) {
	var hexValueLength = 4 // Maximum of two digits menteioned in instructions: \x00 
	indexOfHexValue := strings.Index(input, "\\x")
	
	if indexOfHexValue == -1 {
		return input, 0
	}
	
	var hexValue string
	for index := indexOfHexValue; index < (indexOfHexValue + hexValueLength); index++ {
		hexValue += string(input[index])
	}
	
	// Remove hex value from string and get count of instances it appeared
	output, hexCount := getCountOfAndRemoveSpecifiedSubString(input, hexValue)
	output, moreHexCounts := getCountOfAndRemoveHexCharacters(output)
	
	if moreHexCounts > 0 {
		hexCount += moreHexCounts
	}
	
	return output, hexCount
}

func getCountOfAndRemoveSpecifiedSubString(input string, substring string) (string, int) {
	count := strings.Count(input, substring)
	output := strings.Replace(input, substring, "", -1)
	
	return output, count
}