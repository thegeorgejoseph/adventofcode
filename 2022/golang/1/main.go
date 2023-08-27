package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Advent of Code")
    
    arg1 := os.Args[1]
    if arg1 == "one" {
        One()
    }
    if arg1 == "two" {
        Two()
    }
}
