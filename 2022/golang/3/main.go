package main

import (
    "fmt"
    "os"
)

func Check(err error) {
    fmt.Println("Error: ", err)
    if err != nil {
        panic(err)
    }
}

func main() {
    arg := os.Args[1]
    if (arg == "one"){
        One()
    } else {
        Two()
    }
}
