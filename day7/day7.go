package main

import "fmt"

func main() {	
	var instructions = getInstructionsFromFile()
	fmt.Println(instructions.getCircuitValue("a"))
}