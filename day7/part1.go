package main

/*
--- Day 7: Some Assembly Required ---

This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

123 -> x means that the signal 123 is provided to wire x.
x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.
Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?
*/

import (
	"os"
	"fmt"
	"bufio"	
	"strings"
)

func getInstructionsFromFile() Instructions {
	inFile, _ := os.Open("instructions.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var instructions = Instructions {} 
	for scanner.Scan() {
		instructions = append(instructions, *getInstructionFromString(scanner.Text()))
	}
	return instructions
}

func getInstructionFromString(instructionLine string) *InstructionOperation {
	var instruction = new(InstructionOperation)	
	var operationSegments = strings.Split(instructionLine, operationSeperator)
	var operations = strings.Split(operationSegments[0], " ")	
	// Set the destination that comes after ' -> '
	instruction.destination = operationSegments[1]
	
	switch len(operations) {
		case 1: // e.x. "123" or "lx"
			instruction.operationType = SET
			instruction.argumentOne = operations[0]			
		case 2: // e.x. "NOT bar"
			instruction.operationType = NOT
			instruction.argumentOne = operations[1]			
		case 3: // e.x. "foo LSHIFT 2" or "bar AND foo" 
			instruction.argumentOne = operations[0]
			instruction.argumentTwo = operations[2]
				
			switch operations[1] {
				case "LSHIFT":
					instruction.operationType = LSSHIFT
				case "RSHIFT":
					instruction.operationType = RSHIFT
				case "AND":
					instruction.operationType = AND
				case "OR":
					instruction.operationType = OR
			}
	}	
	
	return instruction
}

func (instructions *Instructions) getInstructionByDestination(name string) (*InstructionOperation, bool) {	
	for _, instruction := range *instructions {
		if instruction.destination == name {
			return &instruction, true
		}
	}
	return nil, false
}

func (instructions *Instructions) getCircuitValue(name string) (uint16, bool) {
	if len(name) <= 0 {
		return 0, false
	}
	
	if value, isNumeric := getNumericValue(name); isNumeric {
		return value, true
	}
	
	instruction, instructionFound := instructions.getInstructionByDestination(name)
	
	if instructionFound == false || instruction == nil {
		return 0, false
	}
	
	if instruction.completed {
		return instruction.value, true
	}
	
	inputValueOne, _ := instructions.getCircuitValue(instruction.argumentOne)	
	inputValueTwo, _ := instructions.getCircuitValue(instruction.argumentTwo)
	
	switch instruction.operationType {
	case SET:
		instruction.value = inputValueOne
	case NOT:
		instruction.value = not(inputValueOne)	
	case RSHIFT:
		instruction.value = rightShift(inputValueOne, inputValueTwo)
	case LSSHIFT:
		instruction.value = leftShift(inputValueOne, inputValueTwo)		
	case AND:
		instruction.value = and(inputValueOne, inputValueTwo)
	case OR:
		instruction.value = or(inputValueOne, inputValueTwo)
	}
	
	instruction.completed = true
	
	fmt.Println("complete", instruction)
	
	return instruction.value, true
}
