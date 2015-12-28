package main

import "fmt"

func main() {
    partOneFilter := RemoveNonNumericCharactersFileFilter {}
    
    fmt.Println("Part 1")
    fmt.Println(getSumOfAccountingFile("accounting.txt", partOneFilter))
    
    fmt.Println("Part 2")
    fmt.Println(getSumOfAccountingFileWithoutRedProperty("accounting.txt"))
}