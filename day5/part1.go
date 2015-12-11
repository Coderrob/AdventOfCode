package main

/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
For example:

ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
jchzalrnumimnmhp is naughty because it has no double letter.
haegwjzuvuyypxyu is naughty because it contains the string xy.
dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?
*/

import (
	"strings"
	"bufio"
	"os"
)

var forbiddenWords = []string { "ab", "cd", "pq", "xy" } 
var vowels = []string { "a", "e", "i", "o", "u" }

func getTotalNiceWordsFromList() int {
	var totalNiceWords int
	inFile, _ := os.Open("naughtyornice.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {
		if (getWhetherStringIsNice(scanner.Text())) {
			totalNiceWords++
		}
	}
	return totalNiceWords	
}

func getWhetherStringIsNice(input string) bool {
	return containsAtLeastThreeVowels(input) && 
		   containsAtLeastOneDoubledCharacter(input) && 
		   doesNotContainForbiddenWords(input)
}

func containsAtLeastThreeVowels(input string) bool {
	var vowelCount int
	for vowelIndex := 0; vowelIndex < len(vowels); vowelIndex++ {
		for index := 0; index < len(input); index++ {
			if (string(input[index]) == vowels[vowelIndex]) {
				vowelCount++
			}
		}
	}
	return vowelCount >= 3
}	

func containsAtLeastOneDoubledCharacter(input string) bool {
	var hasAtLeastOneDuplicatedCharacter = false
	for index := 0; index < len(input); index++ {
		var matchString = string(input[index]) + string(input[index])
		if (strings.Count(input, matchString) > 0) {
			hasAtLeastOneDuplicatedCharacter = true
		}		
	}
	return hasAtLeastOneDuplicatedCharacter
}

func doesNotContainForbiddenWords(input string) bool {
	var doesNotContainForbiddenWords = true	
	for index := 0; index < len(forbiddenWords); index++ {
		if (strings.Contains(input, forbiddenWords[index])) {
			doesNotContainForbiddenWords = false	
		}		
	}
	return doesNotContainForbiddenWords
}