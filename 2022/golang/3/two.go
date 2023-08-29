package main

import (
    "fmt"
    "os"
    "bufio"
)

func getPriority(lines []string) int {
    cache1 := make(map[string]rune)
    cache2 := make(map[string]rune)
    cache3 := make(map[string]rune)
    priority := 0
    line1, line2, line3 := lines[0], lines[1], lines[2]
    CreateMap(line1, cache1)
    CreateMap(line2, cache2)
    CreateMap(line3, cache3)
    for key := range cache1 {
        val, ok2 := cache2[key]
        _, ok3 := cache3[key]
        if ok2 && ok3 {
            keyInt := int(val)
            if keyInt > 96 && keyInt < 123 {
                priority += (keyInt % 97) + 1      
            } else if keyInt > 64 && keyInt < 91 {
                priority += (keyInt % 65) + 27
            }
        }
    }
    return priority
}
func Two() {
    result := 0
    file, err := os.Open("input.txt")
    Check(err)
    scanner := bufio.NewScanner(file)
    lines := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
        if len(lines) == 3 {
            result += getPriority(lines)
            lines = []string{}
        } else {
            continue
        }
    }
    fmt.Println(result)
}
