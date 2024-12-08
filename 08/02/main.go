package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Antenna struct {
	id   string
	x, y int
}

type Antinode struct {
	x, y int
}

type Offset struct {
	x, y int
}

func findAntinodes(ant1 Antenna, ant2 Antenna, maxX int, maxY int) []Antinode {
	// 4. 5. 6.

	/* So there's probably a simple way to find a single offset between them
	   and then usa that to get two "valid" antinodes (in relation to the two
	   antennas), but calculating two offsets works for now
	*/
	var result []Antinode
	antennasOffset1 := Offset{x: ant2.x - ant1.x, y: ant2.y - ant1.y}
	antennasOffset2 := Offset{x: ant1.x - ant2.x, y: ant1.y - ant2.y}

	// From ant1
	possibleAntinode := Antinode{x: ant1.x, y: ant1.y}
	for {
		if possibleAntinode.x < 0 || possibleAntinode.x > maxX ||
			possibleAntinode.y < 0 || possibleAntinode.y > maxY {
			break
		}
		result = append(result, possibleAntinode)
		possibleAntinode.x += antennasOffset1.x
		possibleAntinode.y += antennasOffset1.y
	}

	// From ant2
	possibleAntinode = Antinode{x: ant2.x, y: ant2.y}
	for {
		if possibleAntinode.x <= 0 || possibleAntinode.x >= maxX ||
			possibleAntinode.y <= 0 || possibleAntinode.y >= maxY {
			break
		}
		result = append(result, possibleAntinode)
		possibleAntinode.x += antennasOffset2.x
		possibleAntinode.y += antennasOffset2.y
	}
	return result
}

func main() {

	/*
	   1. Unpack the input into a map[string][]Antenna
	   2. map[Pos]bool to keep antinodes
	   3. Loop through antenna letters; loop through its antennas
	   4. Calculate offset from one to the other
	   5. Propose antinodes that meets the requirements
	   6. check if it's within the map
	   7. If it is, change it to true in the map
	*/

	// 1.
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")
	fileLines := strings.Split(fileStr, "\n")

	antennas := make(map[string][]Antenna)
	for y := range fileLines {
		lineElements := strings.Split(fileLines[y], "")
		for x := range lineElements {
			if lineElements[x] == "." {
				continue
			}
			antenna := Antenna{id: lineElements[x], x: x, y: y}
			antennas[antenna.id] = append(antennas[antenna.id], antenna)
		}
	}

	// 2.
	antinodes := make(map[Antinode]bool)
	for antennaLetter := range antennas {
		for i := range antennas[antennaLetter] {
			currentAntenna := antennas[antennaLetter][i]
			var tmp []Antenna
			tmp = append(tmp, antennas[antennaLetter][:i]...)
			tmp = append(tmp, antennas[antennaLetter][i+1:]...)
			for j := range tmp {
				/* Optimization TODO: right now antennas are checked both ways;
				   once antenna 0 has been checked against all others, antenna
				   7 doesn't have to check against 0 again */
				possibleAntinodes := findAntinodes(currentAntenna, tmp[j],
					len(fileLines[0])-1, len(fileLines)-1)
				for possibleAntinode := range possibleAntinodes {
					antinodes[possibleAntinodes[possibleAntinode]] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
