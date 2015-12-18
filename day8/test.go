package main

import (
	"fmt"
	"os"
	"bufio"
)

func runTestsFromFile() {
	inFile, _ := os.Open("testlist.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	var totalCharacterCount int
	var totalCodeCount int
	for scanner.Scan() {
		 codeCount, characterCount := testCharactersAndCodeCounts(scanner.Text())
		 totalCharacterCount += characterCount
		 totalCodeCount += codeCount
	}	
	fmt.Println(totalCodeCount - totalCharacterCount)
}

func testCharactersAndCodeCounts(testString string) (int, int) {
	var characterCount = getNumberOfCharacters(testString)
	var codeCount = getUpdatedNumberOfCodeCharacters(testString)
	fmt.Println(testString, characterCount, codeCount)
	return codeCount, characterCount
}