package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "slices"
)

func Two() {
    curr := 0
    res := []int{0,0,0}
    file, err := os.Open("input1.txt")
    check(err)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if (line == "") {
            fmt.Println("current", curr)
            if (curr > res[2]) {
                res[2] = curr
                slices.Sort(res)
                slices.Reverse(res)
            }
            curr = 0
            continue
        }
        calories, err := strconv.Atoi(line)
        fmt.Println("Adding Calories: ", calories)
        check(err)
        curr += calories
    }
    if (curr > res[2]) {
        res[2] = curr
        slices.Sort(res)
        slices.Reverse(res)
}
file.Close()
sum := 0
for i := range res {
    sum += res[i]
}
fmt.Println("Total calories: ", sum)
}
