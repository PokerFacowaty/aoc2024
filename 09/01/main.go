package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")

	var current []int
	file := false
	id := -1
	for _, letter := range fileStr {
		file = !file
		if file {
			id += 1
		}

		times, err := strconv.Atoi(string(letter))
		if err != nil {
			log.Fatal(err)
		}

		for j := 0; j < times; j++ {
			if file {
				current = append(current, id)
			} else {
				// -1 to indicate free space since 0 is a valid file Id
				current = append(current, -1)
			}
		}
	}

	i := 0
	j := len(current) - 1
	var compact []int
	checksum := 0
	for i < len(current) && j >= 0 {
		if i > j {
			break
		}
		if current[j] == -1 {
			j--
			continue
		}
		if current[i] == -1 {
			compact = append(compact, current[j])
			checksum += compact[len(compact)-1] * (len(compact) - 1)
			i++
			j--
			continue
		}
		compact = append(compact, current[i])
		checksum += compact[len(compact)-1] * (len(compact) - 1)
		i++
	}
	fmt.Println(checksum)
}
