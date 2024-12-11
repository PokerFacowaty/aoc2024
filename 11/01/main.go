package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
    "slices"
)

func intLen(i int) int {
    if i == 0 {
        return 1
    }

    count := 0
    for i != 0 {
        i /= 10
        count++
    }
    return count
}

func splitNum(i int) (int, int) {
    iStr := strconv.Itoa(i)
    strA, strB := iStr[0:len(iStr) / 2], iStr[len(iStr) / 2:]
    intA, err := strconv.Atoi(strA)
    if err != nil {
        log.Fatal(err)
    }
    intB, err := strconv.Atoi(strB)
    if err != nil {
        log.Fatal(err)
    }
    return intA, intB
}

func main() {
    
	
    fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
   

    stonesStr := strings.Split(strings.Trim(string(fileContent), "\n"), " ")
    var stones []int
    for i := range stonesStr {
        stone, err := strconv.Atoi(stonesStr[i])
        if err != nil {
            log.Fatal(err)
        }

        stones = append(stones, stone)
    }

    for i := 0; i < 25; i++ {
        j := 0
        for {
            if j == len(stones) {
                break
            }
            if stones[j] == 0 {
                stones[j] = 1
                j++
                continue
            }
            digits := intLen(stones[j])
            if digits % 2 == 0 {
                stoneA, stoneB := splitNum(stones[j])
                stones[j] = stoneA
                stones = slices.Insert(stones, j + 1, stoneB)
                j += 2
                continue
            }
            stones[j] *= 2024
            j++
        }
        fmt.Println(len(stones))
    }
    fmt.Println(len(stones))
}
