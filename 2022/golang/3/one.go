package main

import (
    "fmt"
    "bufio"
    "os"
)

func CreateMap(s string, cache map[string]rune) {
    for _, rune := range s {
        _, ok := cache[string(rune)]
        if !ok {
            cache[string(rune)] = rune
        }
    }
}
func One() {
    result := 0
    file, err := os.Open("input.txt")
    Check(err)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        oneCache := make(map[string]rune)
        twoCache := make(map[string]rune)
        line := scanner.Text()
        if line == "" {
            continue
        }
        part1, part2 := line[:len(line)/2], line[len(line)/2:]
        CreateMap(part1, oneCache)
        CreateMap(part2, twoCache)
        for key := range oneCache {
            fmt.Println(key)
            val, ok := twoCache[key]
            fmt.Println(ok)
            if ok{
                //fmt.Println(key, val)
                keyInt := int(val)
                //fmt.Println(keyInt)
                if keyInt > 96 && keyInt < 123 {
                    result += (keyInt % 97) + 1      
                } else if keyInt > 64 && keyInt < 91 {
                    result += (keyInt % 65) + 27
                }
            }
        }
    }
    fmt.Println(result)
}
