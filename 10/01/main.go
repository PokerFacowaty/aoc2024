package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	id, val int // id needed since values repeat
}

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")
	fileLines := strings.Split(fileStr, "\n")
	adjascentList := make(map[Vertex][]Vertex) // optimization: pointers
	var startingOnes []Vertex
	for i := range fileLines {
		line := strings.Split(fileLines[i], "")
		for j := range line {
			id := len(fileLines[i])*i + j
			val, err := strconv.Atoi(line[j])
			if err != nil {
				log.Fatal(err)
			}

			if val == 0 {
				startingOnes = append(startingOnes, Vertex{id, val})
			}

			if i > 0 {
				// Above
				nId := len(fileLines[i])*(i-1) + j
				tmpInt, err := strconv.Atoi(strings.Split(fileLines[i-1], "")[j])
				if err != nil {
					log.Fatal(err)
				}
				adjascentList[Vertex{id, val}] = append(adjascentList[Vertex{id, val}], Vertex{nId, tmpInt})
			}
			if i < len(line)-1 {
				// Below
				nId := len(fileLines[i])*(i+1) + j
				tmpInt, err := strconv.Atoi(strings.Split(fileLines[i+1], "")[j])
				if err != nil {
					log.Fatal(err)
				}
				adjascentList[Vertex{id, val}] = append(adjascentList[Vertex{id, val}], Vertex{nId, tmpInt})
			}
			if j > 0 {
				// Left
				nId := len(fileLines[i])*i + j - 1
				tmpInt, err := strconv.Atoi(strings.Split(fileLines[i], "")[j-1])
				if err != nil {
					log.Fatal(err)
				}
				adjascentList[Vertex{id, val}] = append(adjascentList[Vertex{id, val}], Vertex{nId, tmpInt})
			}
			if j < len(line)-1 {
				// Right
				nId := len(fileLines[i])*i + j + 1
				tmpInt, err := strconv.Atoi(strings.Split(fileLines[i], "")[j+1])
				if err != nil {
					log.Fatal(err)
				}
				adjascentList[Vertex{id, val}] = append(adjascentList[Vertex{id, val}], Vertex{nId, tmpInt})
			}
		}
	}

	scores := 0
	for i := range startingOnes {
		rating := 0
		nines := make(map[Vertex]bool)
		toCheck := []Vertex{} // FIFO list of possbile paths saved for later
		var checking Vertex
		checking = startingOnes[i]
		for {
			fmt.Println("\n\ntoCheck", toCheck)
			fmt.Println("checking", checking)
			if checking.val == 9 {
				fmt.Println("9 found, adding score and continuing")
				nines[checking] = true
				if len(toCheck) > 0 {
					checking, toCheck = toCheck[0], toCheck[1:]
					continue
				}
				break
			}
			neighbors := adjascentList[checking]
			fmt.Println("neighbors:", neighbors)
			var validPathways []Vertex
			for j := range neighbors {
				fmt.Println("neighbor", neighbors[j])
				if neighbors[j].val-checking.val == 1 {
					validPathways = append(validPathways, neighbors[j])
				}
			}
			if len(validPathways) == 0 {
				fmt.Println("No pathways found")
				if len(toCheck) > 0 {
					checking, toCheck = toCheck[0], toCheck[1:]
					continue
				}
				break
			}
			fmt.Println("validPathways:", validPathways)
			checking = validPathways[0]
			toCheck = append(toCheck, validPathways[1:]...)
		}
		scores += len(nines)
	}
	fmt.Println(scores)
}
