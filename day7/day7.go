package main

import "fmt"

func main() {	
	var instructions = getInstructionsFromFile()
	fmt.Println(instructions.getCircuitValue("la"))
	
	for _, instruction := range instructions {
		if instruction.completed {
			fmt.Println(instruction)
		}
	}
}