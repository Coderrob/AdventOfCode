package main

/*
--- Part Two ---

Now find one that starts with six zeroes.
*/

func getHashStartingWithSixZeros(input string) string {
	return getHashStartingWithSpecifiedPrefix(input, "000000")
}