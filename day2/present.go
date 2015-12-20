package main

type Present struct {
    length int
    width int
    height int
}

func (present *Present) getAreaOfSmallestSide() int {
    var areas = []int { 
        (present.length * present.width), 
        (present.width * present.height),
        (present.height * present.length),
    }
	
	return getSmallestSize(areas)
}

func (present *Present) getPaperAmountRequiredToWrap() int {
    var smallestArea = present.getAreaOfSmallestSide()
	var totalSurfaceArea = present.getSurfaceArea()
		
	return smallestArea + totalSurfaceArea
}

func (present *Present) getSurfaceArea() int {
    return (2 * present.length * present.width) + 
		   (2 * present.width * present.height) + 
		   (2 * present.height * present.length)
}

func (present *Present) getRibbonLengthRequiredToWrap() int {	
	var ribbonLength = present.getRibbonLengthOfTwoSmallestSides()
	var bowRibbonLength = present.getRibbonBowLengthRequired()
	
	return ribbonLength + bowRibbonLength
}

func (present *Present) getRibbonLengthOfTwoSmallestSides() int {
    var sizes = getTwoSmallestSizes([]int { 
        present.length, 
        present.width, 
        present.height,
    })
	
	var sizeOne = sizes[0]
	var sizeTwo = sizes[1]
	
	return (sizeOne * 2) + (sizeTwo * 2)
}

func (present *Present) getRibbonBowLengthRequired() int {
    return present.length * present.width * present.height
}