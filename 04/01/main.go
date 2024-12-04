package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func countXmas(board [][]string) int {
	count := 0
	var toCheck [][]string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != "X" {
				continue
			}

			topPossible := i > 2
			botPossible := (len(board) - i) > 3
			leftPossible := j > 2
			rightPossible := (len(board[i]) - j) > 3

			if leftPossible {
				fields := []string{board[i][j], board[i][j-1], board[i][j-2], board[i][j-3]}
				toCheck = append(toCheck, fields)
			}

			if rightPossible {
				fields := []string{board[i][j], board[i][j+1], board[i][j+2], board[i][j+3]}
				toCheck = append(toCheck, fields)
			}

			if topPossible {
				fields := []string{board[i][j], board[i-1][j], board[i-2][j], board[i-3][j]}
				toCheck = append(toCheck, fields)
				if leftPossible {
					fields := []string{board[i][j], board[i-1][j-1], board[i-2][j-2], board[i-3][j-3]}
					toCheck = append(toCheck, fields)
				}
				if rightPossible {
					fields := []string{board[i][j], board[i-1][j+1], board[i-2][j+2], board[i-3][j+3]}
					toCheck = append(toCheck, fields)
				}
			}

			if botPossible {
				fields := []string{board[i][j], board[i+1][j], board[i+2][j], board[i+3][j]}
				toCheck = append(toCheck, fields)
				if leftPossible {
					fields := []string{board[i][j], board[i+1][j-1], board[i+2][j-2], board[i+3][j-3]}
					toCheck = append(toCheck, fields)
				}
				if rightPossible {
					fields := []string{board[i][j], board[i+1][j+1], board[i+2][j+2], board[i+3][j+3]}
					toCheck = append(toCheck, fields)
				}
			}
		}
	}

	for i := 0; i < len(toCheck); i++ {
		if toCheck[i][0] == "X" && toCheck[i][1] == "M" && toCheck[i][2] == "A" && toCheck[i][3] == "S" {
			count++
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
	fmt.Println("XMAS sections:", countXmas(board))
}
