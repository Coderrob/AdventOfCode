package main

import "strconv"

func getLookAndSayResult(input string) string {
	var lookAndSayResult string
	inputLength := len(input)

	for index := 0; index < inputLength; index++ {
		var character = input[index]
		var substring = string(input[index])
		
		for (index + 1) < inputLength && character == input[index + 1] {
			index++
			substring += string(input[index])
		}
		
		lookAndSayResult += getSequenceFromNumericSubString(substring)
	}
		
	return lookAndSayResult
}

func getSequenceFromNumericSubString(input string) string {
	repeatedLength := strconv.Itoa(len(input))
	return repeatedLength + string(input[0])
}
