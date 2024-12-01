package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fileContent, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fileContent), "\n")
	var sliceA []int
	var sliceB []int
	for i := 0; i < len(lines)-1; i++ {
		intA, err := strconv.Atoi(strings.Split(lines[i], "   ")[0])
		if err != nil {
			log.Fatal(err)
		}
		sliceA = append(sliceA, intA)

		intB, err := strconv.Atoi(strings.Split(lines[i], "   ")[1])
		if err != nil {
			log.Fatal(err)
		}
		sliceB = append(sliceB, intB)
	}
	sort.Slice(sliceA, func(i, j int) bool {
		return sliceA[i] < sliceA[j]
	})
	sort.Slice(sliceB, func(i, j int) bool {
		return sliceB[i] < sliceB[j]
	})

	totalDifference := 0
	for i := 0; i < len(sliceA); i++ {
		totalDifference += Abs(sliceA[i] - sliceB[i])
	}
	fmt.Println(totalDifference)

	intCount := make(map[int]int)
	for i := 0; i < len(sliceB); i++ {
		intCount[sliceB[i]]++
	}

	similarityScore := 0
	for i := 0; i < len(sliceA); i++ {
		similarityScore += sliceA[i] * intCount[sliceA[i]]
	}
	fmt.Println(similarityScore)
}
