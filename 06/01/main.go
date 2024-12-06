package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type Guard struct {
	pos       Pos
	direction int // 0 is up and then clockwise until 3
}

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")
	fileLines := strings.Split(fileStr, "\n")

	var guardPos Guard
	obstacles := make(map[Pos]bool)
	for i := range fileLines {
		lineElements := strings.Split(fileLines[i], "")
		for j := range lineElements {
			switch lineElements[j] {
			case "^", "<", ">", "v":
				direction := map[string]int{"^": 0, ">": 1, "v": 2, "<": 3}[lineElements[i]]
				guardPos = Guard{pos: Pos{x: j, y: i}, direction: direction}
			case "#":
				obstacles[Pos{x: j, y: i}] = true
			}
		}
	}

	visited := make(map[Pos]bool)
	for {
		if guardPos.pos.x == -1 || guardPos.pos.x > len(fileLines[0])-1 ||
			guardPos.pos.y == -1 || guardPos.pos.y > len(fileLines)-1 {
			break
		}

		visited[guardPos.pos] = true
		// Move or turn and move the guard
		switch guardPos.direction {
		case 0:
			if obstacles[Pos{x: guardPos.pos.x, y: guardPos.pos.y - 1}] {
				guardPos.direction = 1
				guardPos.pos.x++
				continue
			}
			guardPos.pos.y--
		case 1:
			if obstacles[Pos{x: guardPos.pos.x + 1, y: guardPos.pos.y}] {
				guardPos.direction = 2
				guardPos.pos.y++
				continue
			}
			guardPos.pos.x++
		case 2:
			if obstacles[Pos{x: guardPos.pos.x, y: guardPos.pos.y + 1}] {
				guardPos.direction = 3
				guardPos.pos.x--
				continue
			}
			guardPos.pos.y++
		case 3:
			if obstacles[Pos{x: guardPos.pos.x - 1, y: guardPos.pos.y}] {
				guardPos.direction = 0
				guardPos.pos.y--
				continue
			}
			guardPos.pos.x--
		}
	}
	fmt.Println(len(visited))
}
