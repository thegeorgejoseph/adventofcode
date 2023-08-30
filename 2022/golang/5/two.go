package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func Two() {
    indices := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}
    cache := []map[int]string{}
    boxes := make(map[int][]string)
    file, err := os.Open("input.txt")
    Check(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    for i := 0; i < 9; i ++ {
        scanner.Scan()
        line := scanner.Text()
        lookup := GetValidCharacters(line, indices)
        cache = append(cache, lookup)
    }
    indexLookup := cache[len(cache) - 1]
    cache = cache[:len(cache) - 1]

    for _, tempCache := range cache {
        for tempIndex, value := range tempCache {
            k, err := strconv.Atoi(indexLookup[tempIndex])
            Check(err)
            if _, ok := boxes[k]; !ok {
                boxes[k] = []string{}
            } 
            boxes[k] = append(boxes[k], value)
        }
    }
    for key, values := range boxes {
        boxes[key] = ReversedSlice(values)
    }
    fmt.Println(boxes)
    scanner.Scan()

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        num, err := strconv.Atoi(parts[1])
        Check(err)
        source, err := strconv.Atoi(parts[3])
        Check(err)
        dest, err := strconv.Atoi(parts[5])
        Check(err)
        indexToMove := len(boxes[source]) - num
        newBoxes := boxes[source][indexToMove:]
        boxes[source] = boxes[source][:indexToMove]
        boxes[dest] = append(boxes[dest], newBoxes...)
    }
    
    var result string
    for i := 1; i <= 9; i++ {
        result += boxes[i][len(boxes[i]) - 1]
    }
    fmt.Println(result)
}
