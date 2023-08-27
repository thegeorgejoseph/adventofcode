package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func check(e error) {
    if e != nil { fmt.Println(e)
        panic(e)
    }
}

func One() {
    res, curr := 0, 0
    file, err := os.Open("input1.txt")
    check(err)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if (line == "") {
            fmt.Println("current", curr)
            if (curr > res) {
                res = curr
            }
            //fmt.Println("Current Total", res)
            curr = 0
            continue
        }
        //fmt.Println(line)
        calories, err := strconv.Atoi(line)
        fmt.Println("Adding Calories: ", calories)
        check(err)
        curr += calories
    }
    if (curr > res) {
        res = curr
    }
    file.Close()
    fmt.Println("Total calories: ", res)
}
