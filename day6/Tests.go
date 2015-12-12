package main

import "fmt"

func testOnOffLightsTurningOnAllLights() {
	var onOffLights = generateOnOffLightGrid()
	processCommand(onOffLights, "turn on 0,0 through 999,999")
	fmt.Println(getNumberOfLitOnOffLights(onOffLights))	
}

func testOnOffLightsTurningOnAndOffOneRowOfLights() {
	var onOffLights = generateOnOffLightGrid()
	processCommand(onOffLights, "toggle 0,0 through 999,0")
	fmt.Println(getNumberOfLitOnOffLights(onOffLights))	

	processCommand(onOffLights, "toggle 0,0 through 999,0")
	fmt.Println(getNumberOfLitOnOffLights(onOffLights))		
}

func testOnOffLightsTurnsOffMiddleRowOfLights() {
	var onOffLights = generateOnOffLightGrid()
	processCommand(onOffLights, "turn off 499,499 through 500,500")
	fmt.Println(getNumberOfLitOnOffLights(onOffLights))
}

func testDimmerLightsTurningOnOneLight() {
	var dimmerLights = generateDimmerLightGrid()
	processCommand(dimmerLights, "turn on 0,0 through 0,0")
	fmt.Println(getTotalBrightnessOfDimmerLights(dimmerLights))	
}

func testDimmerLightsToggleAllDimmerLights() {
	var dimmerLights = generateDimmerLightGrid()
	processCommand(dimmerLights, "toggle 0,0 through 999,999")
	fmt.Println(getTotalBrightnessOfDimmerLights(dimmerLights))	
}