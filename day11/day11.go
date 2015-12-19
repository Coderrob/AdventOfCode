package main

import "fmt"

func main() {	
	fmt.Println("Part 1")
    newPassword := generateNewPassword("vzbxkghb")
    fmt.Println(newPassword)
    
    fmt.Println("Part 2")
	fmt.Println(generateNewPassword(newPassword))
}