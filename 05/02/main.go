package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.Trim(string(fileContent), "\n")
	rulesLines := strings.Split(strings.Split(fileStr, "\n\n")[0], "\n")
	pagesLines := strings.Split(strings.Split(fileStr, "\n\n")[1], "\n")

	rules := make(map[string][]string)
	for i := 0; i < len(rulesLines); i++ {
		line := strings.Split(rulesLines[i], "|")
		rules[line[0]] = append(rules[line[0]], line[1])
	}

	totalValid := 0
	totalSorted := 0
	for i := 0; i < len(pagesLines); i++ {
		line := strings.Split(pagesLines[i], ",")
		pages := make(map[string]int)
		for j := 0; j < len(line); j++ {
			// offset by 1 because for non-existent keys there will be 0 returned (int's default value)
			pages[line[j]] = j + 1
		}

		valid := true
		for j := range line {
			needsAfter := rules[line[j]]
			for k := range needsAfter {
				if pages[needsAfter[k]] > 0 && pages[needsAfter[k]] < pages[line[j]] {
					valid = false
					continue
				}
			}
		}

		if valid {
			middleNumberPos := int(len(line) / 2)
			middleNumber, err := strconv.Atoi(line[middleNumberPos])
			if err != nil {
				log.Fatal(err)
			}
			totalValid += middleNumber
		} else {
			sort.Slice(line, func(i, j int) bool {
				// true if i is "bigger" than j
				// i is "bigger" than j if j in not in  in rules[i]
				a := line[i]
				b := line[j]
				for x := range rules[a] {
					if rules[a][x] == b {
						return false
					}
				}
				return true
			})
			middleNumberPos := int(len(line) / 2)
			middleNumber, err := strconv.Atoi(line[middleNumberPos])
			if err != nil {
				log.Fatal(err)
			}
			totalSorted += middleNumber
		}
	}
	fmt.Println(totalSorted)
}
