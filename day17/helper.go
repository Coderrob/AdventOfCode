package main

import (
    "os"
    "bufio"
    "strconv"
)

func loadContainerList(fileName string) []int {
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    containers := []int {}
    for scanner.Scan() {
        size, _ := strconv.Atoi(scanner.Text())
        containers = append(containers, size)
    }
    return containers
}