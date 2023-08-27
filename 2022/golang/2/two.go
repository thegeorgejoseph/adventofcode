package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func newStrategyGuide(player1, player2 string) int {
    // X - lose, Y - draw, Z - win
    var stratscore = map[string]int{"X":0, "Y":3, "Z":6}
    var shapescore = map[string]int{"A":1, "B":2, "C":3}
    score := stratscore[player2]
    var hashmap = map[string]string{"A": "Y", "B": "Z", "C": "X"}
    if (player2 == "Y"){
        score += shapescore[player1]
    } else if (player2 == "Z"){
        score += shapescore[GetCounterPart(hashmap[player1])]
    } else {
        if player1 == "A" {
            score += 3
        } else {
            score += (shapescore[player1] + 2 ) % 3
        }
    }
    return score
}
func Two() {
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
        
        result += newStrategyGuide(player1, player2)
    }
    fmt.Println(result)
    file.Close()
}
