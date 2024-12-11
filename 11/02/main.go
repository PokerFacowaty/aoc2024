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

func splitNum(i int, digits int) (int, int) {
    /*
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
    */

    pow := 1
    for range digits / 2 {
        pow *= 10
    }
    intB := i % pow
    intA := i / pow
    return intA, intB
}

func main() {
    
    blinks := 6
    fileContent, err := os.ReadFile("input_test")
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

    // Calculate the pattern of stone numbers for 0 for all blinks
    testStones := []int{0}
    var pattern [][]int
    for i := 0; i < blinks; i++ {
        j := 0
        var stonesState []int
        // Copying the values
        stonesState = append(stonesState, testStones...)
        pattern = append(pattern, stonesState)
        for {
            if j == len(testStones) {
                break
            }
            if testStones[j] == 0 {
                testStones[j] = 1
                j++
                break
            }
            digits := intLen(testStones[j])
            if digits % 2 == 0 {
                stoneA, stoneB := splitNum(testStones[j], digits)
                testStones[j] = stoneA 
                testStones = slices.Insert(testStones, j + 1, stoneB)
                j += 2
                continue
            }
            testStones[j] *= 2024
            j++
        }
    }
    // fmt.Println(pattern)
    fmt.Println(len(pattern[blinks - 1]))

    // Blink on every stone
    score := 0
    for i := range stones {
        tmpStones := []int{stones[i]}
        for blink := 0; blink < blinks; blink++ {
            j := 0
            for {
                if j == len(tmpStones) {
                    break
                }
                if tmpStones[j] == 0 {
                    score += len(pattern[blinks - 1]) - len(pattern[blink])
                    tmpStones = append(tmpStones[0:j], tmpStones[j+1:]...)
                    continue
                }
                digits := intLen(tmpStones[j])
                if digits % 2 == 0 {
                    stoneA, stoneB := splitNum(tmpStones[j], digits)
                    tmpStones[j] = stoneA
                    tmpStones = slices.Insert(tmpStones, j + 1, stoneB)
                    j += 2
                    continue
                }
                tmpStones[j] *= 2024
                j++
            }
            fmt.Println("After blink:", blink)
            fmt.Println(tmpStones)
        }
        score += len(tmpStones)
    }
    fmt.Println(score)
}
