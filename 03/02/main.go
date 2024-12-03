package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getEnabledInstructions(allInstructions string) string {
	var result []string
	splitByDont := strings.Split(allInstructions, "don't()")
	result = append(result, splitByDont[0])
	for i := 1; i < len(splitByDont); i++ {
		if strings.Contains(splitByDont[i], "do()") {
			splitByDo := strings.Split(splitByDont[i], "do()")
			// result = append(result, strings.Split(splitByDont[i], "do()")[1])
			for i := 1; i < len(splitByDo); i++ {
				result = append(result, splitByDo[i])
			}
		}
	}
	return strings.Join(result, "")
}

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

	validInstructions := getEnabledInstructions(fileStr)

	matches := r.FindAllStringSubmatch(validInstructions, -1)
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
