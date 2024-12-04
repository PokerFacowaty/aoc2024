package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func countXmas(board [][]string) int {
    count := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            // To check
            // Left, right, up, down, left top, right top, left bot, right bot
            if board[i][j] != "X" {
                continue;
            }

            var toCheck [][]*string
            if (i > 2) {
                // Top possible
                toCheck = append(toCheck, [])
            }

        }
    }
    return count
}

func main() {
    fileContent, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }

    fileStr := string(fileContent)
    lines := strings.Split(fileStr, "\n");
    var board [][]string
    for i := 0; i < len(lines) - 1; i++ {
        fmt.Println(lines[i])
        board = append(board, strings.Split(lines[i], ""))
    }
    fmt.Println("XMAS:", countXmas(board))
}
