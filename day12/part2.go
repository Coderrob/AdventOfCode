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

type RemoveRedChildrenFileFilter struct {    
}

func (filter RemoveRedChildrenFileFilter) Clean(fileData []byte) []byte {    
    cleanedFile := []byte {}
    fileLength := len(fileData)
    characterCount := 0
    
    for fileIndex := 0; fileIndex < fileLength; fileIndex++ {        
        var character = fileData[fileIndex]        
        cleanedFile = append(cleanedFile, character)
        characterCount++
        
        // check for instances of a proptery with value of "red" e.x.: ':"red'
        if (fileIndex + 5) < fileLength && 
            character == ':' &&
            fileData[fileIndex + 1] == '"' && 
            fileData[fileIndex + 2] == 'r' &&
            fileData[fileIndex + 3] == 'e' &&
            fileData[fileIndex + 4] == 'd' && 
            fileData[fileIndex + 5] == '"'{
            // determine the index bounds of the object with the property set to "red" for removal
            var openArrays = 0
            var openObjects = 0
            var removeCharacterCount = 0
            // move backwards to find the originating opening { bracket 
            for objectIndex := fileIndex; objectIndex >= 0; objectIndex-- {    
                removeCharacterCount++
                                         
                var subStringCharacter = fileData[objectIndex]
                
                if isOpeningArrayCharacter(subStringCharacter) {
                    openArrays++
                } else if isClosingArrayCharacter(subStringCharacter) {
                    openArrays--
                } else if isOpeningObjectCharacter(subStringCharacter) {
                    openObjects++
                } else if isClosingObjectCharacter(subStringCharacter) {
                    openObjects--
                }
                
                // "red" property was found within an array. No need to filter sub-string.       
                if openArrays == 1 && openObjects == 0 {
                    break
                } 
                
                // "red" property was found within an object. Find end of object and purge the object.
                if openArrays == 0 && openObjects == 1 {
                    for invalidIndex := fileIndex + 5; invalidIndex < fileLength; invalidIndex++ {
                        if isOpeningObjectCharacter(fileData[invalidIndex]) {
                            openObjects++
                        } else if isClosingObjectCharacter(fileData[invalidIndex]) {
                            openObjects--
                        }
                        
                        if openObjects == 0 {
                            fileIndex = invalidIndex
                            break
                        }
                    }
                    
                    characterCount -= removeCharacterCount                    
                    cleanedFile = cleanedFile[0:characterCount]
                    break
                }                
            }
        }   
    }
    
    partOneFilter := RemoveNonNumericCharactersFileFilter {}
    return partOneFilter.Clean(cleanedFile)
}