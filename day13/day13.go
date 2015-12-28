package main

import "fmt"

func main() {
    fmt.Println("Part 1")
    fmt.Println(getHappiestArrangement("attendees.txt"))
    
    fmt.Println("Part 2")
    fmt.Println(getHappiestArrangement("attendeesplusme.txt"))
}