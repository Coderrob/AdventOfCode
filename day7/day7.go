package main

import "fmt"

func main() {	
	var instructions = getInstructionsFromFile()
	fmt.Println("Part 1")
	circuitValue, _ := instructions.getCircuitValue("a")	
	fmt.Println(circuitValue)
	
	fmt.Println("Part 1")	
	instructions.resetValues()	
	instructions.setCircuitValue("b", circuitValue)	
	circuitValueTwo, _ := instructions.getCircuitValue("a")
	fmt.Println(circuitValueTwo)	
}