package main

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

func getOverlap2(lowerA, upperA, lowerB, upperB int) int {
    var largerA, smallerB int

    if (lowerA < lowerB) {
        largerA = upperA
        smallerB = lowerB
    } else {
        largerA = upperB
        smallerB = lowerA
    }
    if (largerA < smallerB) {
        return 0
    }
    return 1
}

func Two() {
    result := 0
    file, err := os.Open("input.txt")
    Check(err)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ",")
        part1, part2 := parts[0], parts[1]
        A, B := strings.Split(part1, "-"), strings.Split(part2, "-")
        lowerA, upperA, lowerB, upperB := ConvertStringToIntegers(A[0]), ConvertStringToIntegers(A[1]), ConvertStringToIntegers(B[0]), ConvertStringToIntegers(B[1])
        result += getOverlap2(lowerA, upperA, lowerB, upperB)
    }
    fmt.Println(result)
}
