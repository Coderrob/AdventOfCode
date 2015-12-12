package main

import "fmt"

func main() {
	fmt.Println("Part 1")	
	var onOffLights = generateOnOffLightGrid()	
	processLightCommandFile(onOffLights)
	fmt.Println(getNumberOfLitOnOffLights(onOffLights))
	
	fmt.Println("Part 2")	
	var dimmerLights = generateDimmerLightGrid()	
	processLightCommandFile(dimmerLights)
	fmt.Println(getTotalBrightnessOfDimmerLights(dimmerLights))
}
