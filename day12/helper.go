package main

import (
    "os"    
    "bufio"
    "strconv"
)

var allowedCharacters = []byte { '[', ']', '{', '}', '-', ',', ':' }

type FileFilter interface {
    Clean(fileData []byte) []byte
}

func getRawAccountingFileData(fileName string) []byte {
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var fileData []byte
    for scanner.Scan() {
		fileData = append(fileData, scanner.Bytes()...)
	}
	return fileData
}

func getCleanedAccountingFile(fileName string, filter FileFilter) []byte {
    accountingFile := getRawAccountingFileData(fileName)
    return filter.Clean(accountingFile)
} 

func getSumOfAccountingFile(fileName string, filter FileFilter) int {
    sumOfAccountingFile := 0
    fileData := getCleanedAccountingFile(fileName, filter)
    fileLength := len(fileData)
    
    for index := 0; index < fileLength; index++ {
        character := fileData[index]
        if isNumericASCIICharacter(character) || character == '-' {
            numberString := string(character)
            for (index + 1) < fileLength && isNumericASCIICharacter(fileData[index + 1]) {
                numberString += string(fileData[index + 1])
                index++
            }
            number, _ := strconv.Atoi(numberString)
            sumOfAccountingFile += number
        }
    }
    return sumOfAccountingFile
}

func isNumericASCIICharacter(input byte) bool {
    return input >= '0' && input <= '9'
}

func isAllowedCharacter(input byte) bool {
    for index := 0; index < len(allowedCharacters); index++ {
        if input == allowedCharacters[index] {
            return true
        }
    }
    return false
}

func getNumericValue(input interface{}) (int, bool) {
    switch input := input.(type) {
        case float64:
            return int(input), true
    }
    return 0, false
}