package main

import "fmt"

func runLookAndSayTests() {
	fmt.Println(getLookAndSayResult("1"))
	fmt.Println(getLookAndSayResult("11"))
	fmt.Println(getLookAndSayResult("21"))
	fmt.Println(getLookAndSayResult("1211"))
	fmt.Println(getLookAndSayResult("111221"))
}