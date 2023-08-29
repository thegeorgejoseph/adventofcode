package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
)

func getOverlap(lowerA, upperA, lowerB, upperB int) int {
    if ((lowerA >= lowerB && upperA <= upperB) || (lowerB >= lowerA && upperB <= upperA)) {
        return 1
    }
    return 0
}

func ConvertStringToIntegers(s string) int {
    i, err := strconv.Atoi(s)
    Check(err)
    return i
}

func One(){
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
        result += getOverlap(lowerA, upperA, lowerB, upperB)
    }
    fmt.Println(result)
}

