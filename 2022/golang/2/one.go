package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func Check(e error) {
    if e != nil {
        fmt.Println("Error: ", e)
    }
}

func GetCounterPart(input string) string {
    var hashmap = map[string]string{"X": "A", "Y": "B", "Z": "C"}
    return hashmap[input]
}

func strategyGuide(player1, player2 string) int {
    var shapescore = map[string]int{"X":1, "Y":2, "Z":3}
    score := shapescore[player2]
    var hashmap = map[string]string{"A": "Y", "B": "Z", "C": "X"}
    if (player1 == GetCounterPart(player2)){
        score += 3
    } else if (player2 == hashmap[player1]){
        score += 6
    }
    return score
}
func One() {
    fmt.Println("One")
    // A - rock, B - paper, C - scissors
    // X - rock, Y - paper, Z - scissors
    
    file, err := os.Open("input.txt")
    Check(err)

    scanner := bufio.NewScanner(file)
    result := 0
    for scanner.Scan() {
        line := scanner.Text()
        if (line == ""){
            continue
        }
        game := strings.Split(line, " ")
        if (len(game) != 2){
            panic("Invalid input")
        }
        player1, player2 := game[0], game[1]
        fmt.Println(player1, player2)
        
        result += strategyGuide(player1, player2)
    }
    fmt.Println(result)
    file.Close()
}
