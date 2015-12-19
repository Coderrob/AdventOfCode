package main

import (
	"strings"
)

// PasswordLengthRequirement is the required length of the Elf security password
const PasswordLengthRequirement = 8

var forbiddenCharacters = []byte { 'i', 'o', 'l' }

func getsWhetherPasswordMeetsSecurityPolicy(input string) bool {
	if meetsRequirement := meetsSecurityLengthRequirement(input); !meetsRequirement {
		return false
	}	
	if hasOnlyLowercaseCharacters := containsOnlyLowercaseCharacters(input); !hasOnlyLowercaseCharacters {
		return false
	}	
	if hasForbiddenCharacters := containsAnyForbiddenCharacters(input); hasForbiddenCharacters {
		return false
	}	
	if hasTwoPairsOfCharacters := containsTwoPairsOfUniqueCharacters(input); !hasTwoPairsOfCharacters {
		return false
	}
	if hasStraightOfCharacters := containsIncreasingStraightOfCharacters(input); !hasStraightOfCharacters {
		return false
	}
	return true
}

func containsOnlyLowercaseCharacters(input string) bool {
	for index := 0; index < len(input); index++ {
		if isLowercaseASCIICharacter(input[index]) == false {
			return false
		}
	}
	return true
}

func isLowercaseASCIICharacter(character byte) bool {
	return character >= 'a' && character <= 'z'
}

func meetsSecurityLengthRequirement(input string) bool {
	return len(input) == PasswordLengthRequirement 
}

func containsAnyForbiddenCharacters(input string) bool {
	for _, forbiddenCharacter := range forbiddenCharacters {
		if strings.ContainsAny(input, string(forbiddenCharacter)) {
			return true
		}
	}
	return false
}

func containsIncreasingStraightOfCharacters(input string) bool {
	inputLength := len(input)
	
	for index := 0; index < inputLength; index++ {
		if (index + 2) < inputLength {
			var character = input[index]
			if (character + 1) == input[index + 1] && (character + 2) == input[index + 2] {				
				return true
			}		
		}
	}
	return false
}

func containsTwoPairsOfUniqueCharacters(input string) bool {
	var pairOfCharacterCount int
	var firstPairsCharacter byte
	
	for index := 0; index < len(input); index++ {
		var character = input[index]
		if strings.Contains(input, string(character) + string(character)) {
			if pairOfCharacterCount == 0 {
				firstPairsCharacter = character
				pairOfCharacterCount++
			} else if firstPairsCharacter != character {
				return true
			}
		}
	}
	return false
}