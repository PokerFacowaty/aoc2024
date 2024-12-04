package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func countXDashMas(board [][]string) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != "A" || i == 0 || i == len(board)-1 || j == 0 || j == len(board)-1 {
				continue
			}

			if (board[i-1][j-1] == "M" && board[i+1][j+1] == "S" || board[i-1][j-1] == "S" && board[i+1][j+1] == "M") && (board[i+1][j-1] == "M" && board[i-1][j+1] == "S" || board[i+1][j-1] == "S" && board[i-1][j+1] == "M") {
				count++
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
	lines := strings.Split(fileStr, "\n")
	var board [][]string
	for i := 0; i < len(lines)-1; i++ {
		board = append(board, strings.Split(lines[i], ""))
	}
	fmt.Println("X-MAS sections: ", countXDashMas(board))
}
