package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Question 2")
    arg1 := os.Args[1]
    if (arg1 == "one"){
        One()
    }
    if (arg1 == "two"){
        Two()
    }
}
