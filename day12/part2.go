package main

/*
--- Part Two ---

Uh oh - the Accounting-Elves have realized that they double-counted everything red.

Ignore any object (and all of its children) which has any property with the value "red". Do this only for objects ({...}), not arrays ([...]).

[1,2,3] still has a sum of 6.
[1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is ignored.
{"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is ignored.
[1,"red",5] has a sum of 6, because "red" in an array has no effect.
*/

import "fmt"

type RemoveRedChildrenFileFilter struct {    
}

func (filter RemoveRedChildrenFileFilter) Clean(fileData []byte) []byte {    
    cleanedFile := []byte {}
    fileLength := len(fileData)
    openedSeperators := 0
    redWordEncountered := false
    
    for index := 0; index < fileLength; index++ {        
        character := fileData[index]
        
        if isOpeningObjectCharacter(character) == false {
            cleanedFile = append(cleanedFile, character)
            continue
        }        
        
        openedSeperators++
        for objectIndex := index + 1; objectIndex < fileLength; objectIndex++ {
            if (objectIndex + 2) < fileLength && fileData[objectIndex] == 'r' && fileData[objectIndex + 1] == 'e' && fileData[objectIndex + 2] == 'd' {            
                redWordEncountered = true
            }
            
            if isOpeningObjectCharacter(fileData[objectIndex]) {                
                if redWordEncountered {
                    openedSeperators++
                } else {
                    openedSeperators = 0
                    break
                }
            } else if isClosingObjectCharacter(fileData[objectIndex]) {
                openedSeperators--
                
                if openedSeperators == 0 && redWordEncountered {
                    redWordEncountered = false
                    index = objectIndex + 1                    
                }
            }
        }
    }
    
    fmt.Println(string(cleanedFile))
    
    partOneFilter := RemoveNonNumericCharactersFileFilter {}
    return partOneFilter.Clean(cleanedFile)
}

func isOpeningObjectCharacter(character byte) bool {
    return character == '{'
}

func isClosingObjectCharacter(character byte) bool {
    return character == '}'
}