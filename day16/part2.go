package main

/*
--- Part Two ---

As you're about to send the thank you note, something in the MFCSAM's instructions catches your eye. Apparently, it has an outdated retroencabulator, and so the output from the machine isn't exact values - some of them indicate ranges.

In particular, the cats and trees readings indicates that there are greater than that many (due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians and goldfish readings indicate that there are fewer than that many (due to the modial interaction of magnetoreluctance).

What is the number of the real Aunt Sue?
*/

func findAuntSueIDByPropertyRanges(auntSueList []AuntSue, properties map[string]int) (int, bool) {
    for _, auntSue := range auntSueList {
        hasAllProperties := true        
        for property := range auntSue.properites {
            hasProperty := false
            
            if property == "cats" || property == "trees" {
                hasProperty = properties[property] < auntSue.properites[property]
            } else if property == "pomeranians" || property == "goldfish" {
                hasProperty = properties[property] > auntSue.properites[property]
            } else {
                hasProperty = properties[property] == auntSue.properites[property]
            }
            
            if hasProperty == false {
                hasAllProperties = false
            }
        }
        
        if hasAllProperties {
            return auntSue.id, true
        }
    }
    return 0, false
}