package main

import (
    "os"
    "bufio"
)

var allowedCharacters = []byte { '[', ']', '{', '}', '-', ',', ':' }

type FileFilter interface {
    Clean(fileData []byte) []byte
}

func getRawAccountingFileData() []byte {
    file, _ := os.Open("accounting.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var fileData []byte
    for scanner.Scan() {
		fileData = append(fileData, scanner.Bytes()...)
	}
	return fileData
}

func getCleanedAccountingFile(filter FileFilter) []byte {
    accountingFile := getRawAccountingFileData()
    return filter.Clean(accountingFile)
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