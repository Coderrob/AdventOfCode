package main

/*
--- Part Two ---

The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
How many total feet of ribbon should they order?
*/

// l w h
import (
	"bufio"
	"os"
)

func getTotalRibbonRequired() int {
	var totalRibbonRequired int
	inFile, _ := os.Open("ribbon.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
	for scanner.Scan() {		
		totalRibbonRequired += getRibbonRequired(scanner.Text())
	}
	return totalRibbonRequired	
}

func getRibbonRequired(str string) int {
	var dimeneions = getDimentionsFromString(str)
	
	var length = dimeneions[0]
	var width = dimeneions[1]
	var height = dimeneions[2]
	
	var ribbonLength = getRibbonLengthByDimension(length, width, height)
	var bowRibbonLength = getBowRibbonLengthByDimension(length, width, height)
	
	return ribbonLength + bowRibbonLength
}

func getRibbonLengthByDimension(length int, width int, height int) int {
	var sizes = getTwoSmallestSizes([]int { length, width, height })
	
	var sizeOne = sizes[0]
	var sizeTwo = sizes[1]
	
	return (sizeOne * 2) + (sizeTwo * 2)
}

func getBowRibbonLengthByDimension(length int, width int, height int) int {
	return length * width * height
}