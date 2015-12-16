package main

import "fmt"

func main() {	
	var instructions = getInstructionsFromFile()
	fmt.Println("=========================")
	fmt.Println(instructions.getCircuitValue("a"))	
	fmt.Println(instructions.getCircuitValue("a"))	
}