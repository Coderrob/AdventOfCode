package main

import "fmt"

func runFilterTests() {    
    partTwoFilter := RemoveRedChildrenFileFilter {}
    fmt.Println(getSumOfAccountingFile("examples.txt", partTwoFilter) == 16)
    fmt.Println(getSumOfAccountingFile("testaccounting.txt", partTwoFilter))
}