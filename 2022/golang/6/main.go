package main

import (
    "fmt"
    "os"
    "bufio"
)

func Check(e error) {
    if e != nil {
        fmt.Println("Error: ", e)
        panic(e)
    }
}

func ReadInput() string {
    file, err := os.Open("input.txt")
    Check(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    return scanner.Text()
}

func main() {
    args := os.Args[1]
    if args == "one" {
        fmt.Println(One(ReadInput()))
    } else {
        Two(ReadInput())
    }
}
