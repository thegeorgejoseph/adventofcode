package main

import (
    "fmt"
    "os"
)

func Check(err error){

    if err != nil {
        fmt.Println("Error: ", err)
        panic(err)
    }
}

func main() {
    
    arg := os.Args[1]
    if arg == "one" {
        One()
    } else {
        Two()
    }
}
