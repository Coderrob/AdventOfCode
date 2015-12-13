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
	"strconv"
	"strings"
)

var operationSeperator = " -> "
var circuits = map[string]uint16 {}
var circuitNames = []string {} 
var instructions = []string {}

func loadCircuitInstructionFile() {
	inFile, _ := os.Open("instructions.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
}

func processInstructionsUntilAllCircuitsValuesAreFound() {
	var instructionsComplete bool
	for instructionsComplete == false {
		instructionsComplete = true
		for index := 0; index < len(instructions); index++ {
			if couldPerformInstuctionLine(instructions[index]) != true {				
				instructionsComplete = false
			}			
		}
	}
}

func doesCircuitExist(circuit string) bool {
    for _, value := range circuitNames {
        if value == circuit {
            return true
        }
    }
    return false
}

func addCircuit(circuit string) {
	if doesCircuitExist(circuit) == false {
		circuitNames = append(circuitNames, circuit)
	}
}

func couldPerformInstuctionLine(instructionLine string) bool {
	var operations = strings.Split(instructionLine, operationSeperator)
	var leftOperation = operations[0]	
	var destinationCircuit = operations[1]
	var leftOperations = strings.Split(leftOperation, " ")
	
	switch len(leftOperations) {
		case 1: // e.x. "123"
			initialValue, _ := strconv.ParseUint(leftOperations[0], 0, 16)	
			circuits[destinationCircuit] = uint16(initialValue)
			addCircuit(destinationCircuit)
			
		case 2: // e.x. "NOT bar"
			var circuit = leftOperations[1]
			
			if doesCircuitExist(circuit) == false {
				return false
			}
			
			circuits[destinationCircuit] = notOperation(circuits[circuit])			
			addCircuit(destinationCircuit)	
			
		case 3: // e.x. "foo LSHIFT 2" or "bar AND foo"
			var circuitOne = leftOperations[0]
			var operationType = leftOperations[1]
			
			// No sense in continuing if the first circuit wasn't added yet
			if doesCircuitExist(circuitOne) == false {
				return false	
			}
			
			if operationType == "LSHIFT" || operationType == "RSHIFT" {							
				shiftValue, _ := strconv.ParseUint(leftOperations[2], 0, 16)								
				switch operationType {
					case "LSHIFT":
						circuits[destinationCircuit] = leftShiftOperation(circuits[circuitOne], uint16(shiftValue))
					case "RSHIFT":
						circuits[destinationCircuit] = rightShiftOperation(circuits[circuitOne], uint16(shiftValue))
				}				
				addCircuit(destinationCircuit)				
			} else if operationType == "ADD" || operationType == "OR" {			
				var circuitTwo = leftOperations[2]										
				
				if doesCircuitExist(circuitTwo) == false {
					return false
				}
				
				switch operationType {
					case "AND":
						circuits[destinationCircuit] = andOperation(circuits[circuitOne], circuits[circuitTwo])
					case "OR":
						circuits[destinationCircuit] = orOperation(circuits[circuitOne], circuits[circuitTwo])
				}			
				addCircuit(destinationCircuit)
			}
	}

	return true
}

func writeOutCircuitValues(circuits map[string]uint16) {
	fmt.Println("-------")
	for key, value := range circuits {
		fmt.Println(key, value)
	}
}

func writeOutCircuitNames(circuits []string) {
	fmt.Println("-------")
	for key, value := range circuits {
		fmt.Println(key, value)
	}
}

func andOperation(circuitOne uint16, circuitTwo uint16) uint16 {
	return circuitOne & circuitTwo
}

func orOperation(circuitOne uint16, circuitTwo uint16) uint16 {
	return circuitOne | circuitTwo
}

func notOperation(circuitValue uint16) uint16 {
	return ^circuitValue
}

func leftShiftOperation(circuitValue uint16, shiftValue uint16) uint16 {
	return circuitValue << shiftValue
}

func rightShiftOperation(circuitValue uint16, shiftValue uint16) uint16 {
	return circuitValue >> shiftValue
}