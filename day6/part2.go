package main

/*
You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.
*/

import (
	"fmt"
)

func generateDimmerLightGrid() dimmerLights {
	lights := dimmerLights { lightGrid: make(map[string]int) }
	for index := 0; index < 1000; index++ {
		var coordinate = fmt.Sprintf("%d,%d", index, index) 
		lights.lightGrid[coordinate] = 0
	}
	return lights
}

func getTotalBrightnessOfDimmerLights(dimmmerLights dimmerLights) int {
	var totalLightBrightnessCound int
	for xCoordinate := 0; xCoordinate < 1000; xCoordinate++ {
		for yCoordinate := 0; yCoordinate < 1000; yCoordinate++ {
			var coordinate = fmt.Sprintf("%d,%d", xCoordinate, yCoordinate)
			totalLightBrightnessCound += dimmmerLights.lightGrid[coordinate]
		}
	}
	return totalLightBrightnessCound
}