package main

import "fmt"

func main() {
	fmt.Println(getHashStartingWithFiveZeros("abcdef"))	
	
	fmt.Println(getHashStartingWithFiveZeros("pqrstuv"))	
	
	fmt.Println(getHashStartingWithFiveZeros("yzbqklnj"))
	
	fmt.Println(getHashStartingWithSixZeros("yzbqklnj"))
}