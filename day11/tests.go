package main

import "fmt"

func runPasswordValidationTests() {
	fmt.Println(getsWhetherPasswordMeetsSecurityPolicy("hijklmmn") == false)
	fmt.Println(getsWhetherPasswordMeetsSecurityPolicy("abbceffg") == false)
	fmt.Println(getsWhetherPasswordMeetsSecurityPolicy("abbcegjk") == false)

	fmt.Println(meetsSecurityLengthRequirement("a") == false)
	fmt.Println(meetsSecurityLengthRequirement("1") == false)
	fmt.Println(meetsSecurityLengthRequirement("aaaaaaaa") == true)

	fmt.Println(isLowercaseASCIICharacter('0') == false)
	fmt.Println(isLowercaseASCIICharacter('z' + 1) == false)
	fmt.Println(isLowercaseASCIICharacter('a') == true)
	fmt.Println(isLowercaseASCIICharacter('z') == true)	
	
	fmt.Println(containsOnlyLowercaseCharacters("A") == false)
	fmt.Println(containsOnlyLowercaseCharacters("a") == true)

	fmt.Println(containsAnyForbiddenCharacters("abcd") == false)
	fmt.Println(containsAnyForbiddenCharacters("abcdi") == true)
	fmt.Println(containsAnyForbiddenCharacters("abocd") == true)
	fmt.Println(containsAnyForbiddenCharacters("albcd") == true)
	
	fmt.Println(containsTwoPairsOfUniqueCharacters("abccd") == false)
	fmt.Println(containsTwoPairsOfUniqueCharacters("aabccd") == true)
	fmt.Println(containsTwoPairsOfUniqueCharacters("aadbccdd") == true)

	fmt.Println(containsIncreasingStraightOfCharacters("abecd") == false)
	fmt.Println(containsIncreasingStraightOfCharacters("abccd") == true)
	fmt.Println(containsIncreasingStraightOfCharacters("pabcd") == true)
	fmt.Println(containsIncreasingStraightOfCharacters("pabcf") == true)

	//fmt.Println(generateNewPassword("abcdefgh") == "abcdffaa") // abcdffaa
	//fmt.Println(generateNewPassword("ghijklmn") == "ghjaabcc") // ghjaabcc
}