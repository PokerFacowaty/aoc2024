package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := string(fileContent)
	// NTS: Using backticks to get a raw string
	r, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(fileStr, -1)
	total := 0
	for i := 0; i < len(matches); i++ {
		intA, err := strconv.Atoi(matches[i][1])
		if err != nil {
			log.Fatal(err)
		}

		intB, err := strconv.Atoi(matches[i][2])
		if err != nil {
			log.Fatal(err)
		}
		total += intA * intB
	}
	fmt.Println(total)
}
