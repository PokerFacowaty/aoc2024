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
	rulesLines := strings.Split(strings.Split(fileStr, "\n\n")[0], "\n")
	pagesLines := strings.Split(strings.Split(fileStr, "\n\n")[1], "\n")

	rules := make(map[string][]string)
	for i := 0; i < len(rulesLines); i++ {
		line := strings.Split(rulesLines[i], "|")
		rules[line[0]] = append(rules[line[0]], line[1])
	}

	total := 0
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
			total += middleNumber
		}
	}
	fmt.Println(total)
}
