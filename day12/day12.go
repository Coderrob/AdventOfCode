package main

import "fmt"

func main() {
    partOneFilter := RemoveNonNumericCharactersFileFilter {}
    
    fmt.Println("Part 1")
    fmt.Println(getSumOfAccountingFile("accounting.txt", partOneFilter))
    
    partTwoFilter := RemoveRedChildrenFileFilter {}
    
    fmt.Println("Part 2")
    fmt.Println(getSumOfAccountingFile("accounting.txt", partTwoFilter))
}