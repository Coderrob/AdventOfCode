package main

import (
	"fmt"
	"crypto/md5"
	"strings"
	"strconv"
)

func getMD5Hash(input string) string {
	data := []byte(input)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func getHashStartingWithSpecifiedPrefix(input string, prefix string) string {
	var number int
	var key string
	var hash string
	
	for (strings.HasPrefix(hash, prefix) != true) {
		key = strconv.Itoa(number)
		hash = getMD5Hash(input + key)
		number++		
	}
	
	return input + key
}