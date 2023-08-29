package main

import (
    "fmt"
    "os"
)

func Check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    fmt.Println("Day 4")
    arg := os.Args[1]
    if arg == "one" {
        One()
    } else {
        Two()
    }
}
