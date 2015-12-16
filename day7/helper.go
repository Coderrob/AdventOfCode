package main

import "strconv"

var operationSeperator = " -> "

const (
	SET = 0
	NOT = 1
	LSSHIFT = 2
	RSHIFT = 3
	AND = 4
	OR = 5
)

type Instructions []InstructionOperation 

type InstructionOperation struct {	
	operationType int
	completed bool
	value uint16
	destination string
	argumentOne string
	argumentTwo string
}

func getNumericValue(value string) (uint16, bool) {
	initialValue, err := strconv.ParseUint(value, 0, 16)
	if err == nil {
		return uint16(initialValue), true
	} 
	return 0, false
}

func and(valueOne uint16, valueTwo uint16) uint16 {
	return valueOne & valueTwo
}

func or(valueOne uint16, valueTwo uint16) uint16 {
	return valueOne | valueTwo
}

func not(value uint16) uint16 {
	return ^value
}

func leftShift(valueOne uint16, valueTwo uint16) uint16 {
	return valueOne << valueTwo
}

func rightShift(valueOne uint16, valueTwo uint16) uint16 {
	return valueOne >> valueTwo
}