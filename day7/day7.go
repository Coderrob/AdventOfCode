package main

func main() {
	loadCircuitInstructionFile()
	processInstructionsUntilAllCircuitsValuesAreFound()
	writeOutCircuitValues(circuits)
}