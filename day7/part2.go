package main

/*
--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?
*/

import "fmt"

func (instructions Instructions) resetValues() {
	for index := 0; index < len(instructions); index++ {
		instructions[index].completed = false
		instructions[index].value = 0
	}
}

func (instructions Instructions) setCircuitValue(name string, value uint16) {
	instruction, index, _ := instructions.getInstructionByDestination(name)
	fmt.Println(instruction)
	
	instruction.operationType = SET		
	instruction.argumentOne = string(value)	
	instruction.completed = true
	instruction.value = value
	
	instructions[index] = instruction
}