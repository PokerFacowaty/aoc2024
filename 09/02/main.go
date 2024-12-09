package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Space struct {
	id   int // -1 for free space
	size int
}

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")

	var current []Space
	file := false
	id := -1
	for _, letter := range fileStr {
		file = !file
		if file {
			id += 1
		}

		size, err := strconv.Atoi(string(letter))
		if err != nil {
			log.Fatal(err)
		}

		if file {
			current = append(current, Space{id: id, size: size})
		} else {
			// -1 to indicate free space since 0 is a valid file Id
			current = append(current, Space{id: -1, size: size})
		}
	}

	// Slow idea that later got patches to patches, but worked
	idsMovedAlready := make(map[int]bool)
	for i := len(current) - 1; i > 0-1; i-- {
		if current[i].id == -1 {
			// Free space
			continue
		}

		if idsMovedAlready[current[i].id] {
			continue
		}

		for j := 0; j < len(current); j++ {
			if j > i {
				break
			}
			// Finding a suitable space
			if current[j].id == -1 {
				if current[j].size == current[i].size {
					current[j].id = current[i].id
					idsMovedAlready[current[i].id] = true
					current[i].id = -1
					break
				}
				if current[j].size > current[i].size {
					current[j].size -= current[i].size
					current = slices.Insert(current, j, Space{id: current[i].id, size: current[i].size})
					idsMovedAlready[current[i].id] = true
					i++
					current[i].id = -1
					break
				}
			}
		}
	}
	checksum := 0
	multiplier := 0
	slIndex := 0
	for slIndex < len(current) {
		if current[slIndex].size == 0 {
			slIndex++
			continue
		}
		for x := 0; x < current[slIndex].size; x++ {
			if current[slIndex].id > -1 {
				checksum += current[slIndex].id * multiplier
			}
			multiplier++
		}
		slIndex++
	}
	fmt.Println(checksum)
}
