package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

var commandPrefixes = []string { "toggle ", "turn on ", "turn off " }
var commandCoordinateSplitWord = " through "

var toggleLight byte = 0
var turnOnLight byte = 1
var turnOffLight byte = 2

type lights interface {	
	toggle(coordinate string)
	turnOn(coordinate string)
	turnOff(coordinate string)
}

type dimmerLights struct {
	lightGrid map[string]int
}

type onOffLights struct {
	lightGrid map[string]bool
}

func (dimmerLights dimmerLights) turnOn(coordinate string) {
	dimmerLights.lightGrid[coordinate]++
}

func (dimmerLights dimmerLights) turnOff(coordinate string) {
	var currentValue = dimmerLights.lightGrid[coordinate]
	if (currentValue <= 0) {
		currentValue = 0
	} else {
		currentValue--
	}
	dimmerLights.lightGrid[coordinate] = currentValue
}

func (dimmerLights dimmerLights) toggle(coordinate string) {
	dimmerLights.lightGrid[coordinate] += 2
}

func (onOffLights onOffLights) turnOn(coordinate string) {
	onOffLights.lightGrid[coordinate] = true
}

func (onOffLights onOffLights) turnOff(coordinate string) {
	onOffLights.lightGrid[coordinate] = false
}

func (onOffLights onOffLights) toggle(coordinate string) {
	onOffLights.lightGrid[coordinate] = !onOffLights.lightGrid[coordinate]
}

func processLightCommandFile(lights lights) {
	inFile, _ := os.Open("lightinstructions.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {
		processCommand(lights, scanner.Text())
	}
}

func processCommand(lights lights, command string) {
	if (strings.HasPrefix(command, commandPrefixes[toggleLight])) {
		command = strings.TrimPrefix(command, commandPrefixes[toggleLight])
		var ranges = strings.Split(command, commandCoordinateSplitWord)
		toggleLights(lights, ranges[0], ranges[1])
	} else if (strings.HasPrefix(command, commandPrefixes[turnOnLight])) {
		command = strings.TrimPrefix(command, commandPrefixes[turnOnLight])
		var ranges = strings.Split(command, commandCoordinateSplitWord)
		turnLightsOn(lights, ranges[0], ranges[1])
	} else if (strings.HasPrefix(command, commandPrefixes[turnOffLight])) {
		command = strings.TrimPrefix(command, commandPrefixes[turnOffLight])
		var ranges = strings.Split(command, commandCoordinateSplitWord)
		turnLightsOff(lights, ranges[0], ranges[1])
	}
}

func toggleLights(lights lights, startRange string, endRange string) {
	handleLightChangeByType(lights, startRange, endRange, toggleLight)
}

func turnLightsOn(lights lights, startRange string, endRange string) {	
	handleLightChangeByType(lights, startRange, endRange, turnOnLight)
}

func turnLightsOff(lights lights, startRange string, endRange string) {
	handleLightChangeByType(lights, startRange, endRange, turnOffLight)
}

func handleLightChangeByType(lights lights, startRange string, endRange string, changeType byte) {	
	var startCoordiantes = strings.Split(startRange, ",")
	var endCoordinates = strings.Split(endRange, ",")
	
	xStartCoordinate, _ := strconv.Atoi(startCoordiantes[0])
	yStartCoordinate, _ := strconv.Atoi(startCoordiantes[1])
	
	xEndCoordinate, _ := strconv.Atoi(endCoordinates[0])
	yEndCoordiante, _ := strconv.Atoi(endCoordinates[1])
	
	for xCoordinate := xStartCoordinate; xCoordinate <= xEndCoordinate; xCoordinate++ {
		for yCoordinate := yStartCoordinate; yCoordinate <= yEndCoordiante; yCoordinate++ {
			var coordinate = fmt.Sprintf("%d,%d", xCoordinate, yCoordinate)
			switch changeType {
				case toggleLight:
					lights.toggle(coordinate)			 
				case turnOnLight:
					lights.turnOn(coordinate)
				case turnOffLight:
					lights.turnOff(coordinate)
			}
		}
	}
}