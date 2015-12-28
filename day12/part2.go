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

import (
    "encoding/json"
)

func getSumOfAccountingFileWithoutRedProperty(fileName string) int {
    fileData := getRawAccountingFileData(fileName)
    array := []interface{} {}
    
    if err := json.Unmarshal(fileData, &array); err != nil {
        return 0
    }
    
    return processArray(array)
}

func processArray(array []interface{}) int {
	total := 0
	for _, arrayValue := range array {
		if value, ok := getNumericValue(arrayValue); ok {
			total += value
            continue
        }
        
        if array, ok := arrayValue.([]interface{}); ok {
			total += processArray(array)
            continue
		} 
        
        if objectMap, ok := arrayValue.(map[string]interface{}); ok {
			total += processMap(objectMap)
        }
	}
	return total
}

func processObject(object interface{}) int {
	if value, ok := getNumericValue(object); ok {
        return value
    }
    
	if array, ok := object.([]interface{}); ok {
		return processArray(array)
	}

	if objectMap, ok := object.(map[string]interface{}); ok {
		return processMap(objectMap)
	}
    
	return 0
}

func processMap(object map[string]interface{}) int {
	total := 0
	for _, field := range object {
        if value, isString := field.(string); isString {
			if value == "red" {
				return 0
			}
			continue
		}
        
		if value, ok := getNumericValue(field); ok {
            total += value
			continue
        }

		if array, ok := field.([]interface{}); ok {
			total += processArray(array)
			continue
		}
                
		if object, ok := field.(map[string]interface{}); ok {
			total += processObject(object)
			continue
		}
	}
	return total
}