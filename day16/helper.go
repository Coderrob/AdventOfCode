package main

import (
    "os"    
    "bufio"
    "strconv"
    "strings"
)

type AuntSue struct {
    id int 
    properites map[string]int 
}

func loadAuntSueList(fileName string) []AuntSue {
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    auntSueList := []AuntSue {}
    for scanner.Scan() {
        auntSueList = append(auntSueList, getAuntSueFromString(scanner.Text()))
    }
    return auntSueList
}

func getAuntSueFromString(input string) AuntSue {
    input = strings.Replace(input, "Sue ", "", -1)
    input = strings.Replace(input, ",", ":", -1)
    
    partialInputs := strings.Split(input, ": ")
    splitLength := len(partialInputs)
    
    auntSue := AuntSue {}
    auntSue.properites = map[string]int {}
    
    for index := 0; index < splitLength; index++ {
        if index == 0 {
            auntID, _ := strconv.Atoi(partialInputs[index])
            auntSue.id = auntID
            continue
        }
        
        if index % 2 != 0 && (index + 1) < splitLength {
            value, _ := strconv.Atoi(partialInputs[index + 1])
            auntSue.properites[partialInputs[index]] = value
            index++
        }
    }
    return auntSue
}