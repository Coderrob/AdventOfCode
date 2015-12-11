package main

/*
--- Part Two ---

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
For example:

qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.
How many strings are nice under these new rules?
*/

import (
	"strings"
	"bufio"
	"os"
)

func getTotalNiceWordsFromNewRulesList() int {
	var totalNiceWords int
	inFile, _ := os.Open("naughtyornice.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {
		if (getWhetherStringIsNiceWithNewRules(scanner.Text())) {
			totalNiceWords++
		}
	}
	return totalNiceWords	
}

func getWhetherStringIsNiceWithNewRules(input string) bool {
	return containsCharacterSetThatDoNotOverlap(input) &&
		   containsOneCharacterThatRepeatsWithDifferentMiddleCharacter(input)
}

func containsCharacterSetThatDoNotOverlap(input string) bool {
	var containsTwoRepeatingCharacters bool
	stringLength := len(input)
	for index := 0; index < stringLength; index++ {
		if ((index + 1) < stringLength) {
			var characterSet = string(input[index]) + string(input[index + 1])			
			if (strings.Count(input, characterSet) > 1) {
				containsTwoRepeatingCharacters = true
			}
		} 
		
		if ((index - 1) > -1) {			
			var characterSet = string(input[index - 1]) + string(input[index])			
			if (strings.Count(input, characterSet) > 1) {
				containsTwoRepeatingCharacters = true
			}
		}		
	}
	return containsTwoRepeatingCharacters
}

func containsOneCharacterThatRepeatsWithDifferentMiddleCharacter(input string) bool {
	var containsRepeatedCharacter bool
	stringLength := len(input)
	for index := 0; index < stringLength; index++ {		
		var character = input[index]
		if ((index + 2) < stringLength) {
			if (character != input[index + 1] && character == input[index + 2]) {
				containsRepeatedCharacter = true
			}
		}
	}
	return containsRepeatedCharacter
}