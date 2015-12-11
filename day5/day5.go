package main

import "fmt"

func main() {
	//fmt.Println("Part 1")
	//fmt.Println(getTotalNiceWordsFromList())
	
	fmt.Println("Part 2")
	fmt.Println(getTotalNiceWordsFromNewRulesList())
}

func part1Tests() {
	fmt.Println(getWhetherStringIsNice("aaa"))
	fmt.Println(getWhetherStringIsNice("jchzalrnumimnmhp"))
	fmt.Println(getWhetherStringIsNice("haegwjzuvuyypxyu"))
	fmt.Println(getWhetherStringIsNice("dvszwmarrgswjxmb"))
}

func part2Tests() {
	fmt.Println(getWhetherStringIsNiceWithNewRules("qjhvhtzxzqqjkmpb")) //nice
	fmt.Println(getWhetherStringIsNiceWithNewRules("xxyxx")) //nice
	fmt.Println(getWhetherStringIsNiceWithNewRules("uurcxstgmygtbstg")) //naught
	fmt.Println(getWhetherStringIsNiceWithNewRules("ieodomkazucvgmuy")) //naught
	
	fmt.Println(containsCharacterSetThatDoNotOverlap("xyxy")) // true
	fmt.Println(containsCharacterSetThatDoNotOverlap("aabcdefgaa")) // true
	fmt.Println(containsCharacterSetThatDoNotOverlap("aaa")) // false
	
	fmt.Println(containsOneCharacterThatRepeatsWithDifferentMiddleCharacter("xyx")) // true
	fmt.Println(containsOneCharacterThatRepeatsWithDifferentMiddleCharacter("abcdefeghi")) // true
	fmt.Println(containsOneCharacterThatRepeatsWithDifferentMiddleCharacter("aaa")) // true
}
