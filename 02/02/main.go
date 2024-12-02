package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func isSafeWithoutOne(lineNums []int) bool {
    for i := 0; i < len(lineNums); i++ {
        var temp []int
        temp = append(temp, lineNums[:i]...)
        temp = append(temp, lineNums[i + 1:]...)
        if isSafe(temp) {
            return true
        }
    }
    return false
}

func isSafe(lineNums []int) bool {
    increasing := false
    if (lineNums[1] - lineNums[0]) > 0 {
        increasing = true
    }

    for i := 1; i < len(lineNums); i++ {
        if lineNums[i] == lineNums[i - 1] {
            return false
        } else if increasing && ((lineNums[i] - lineNums[i - 1]) < 0) {
            return false
        } else if !increasing && ((lineNums[i] - lineNums[i - 1]) > 0) {
            return false
        } else if Abs(lineNums[i] - lineNums[i - 1]) > 3 {
            return false
        }
    }
    return true
}

func main() {
    fileContent, err := os.ReadFile("input")

    if err != nil {
        log.Fatal(err)
    }

    lines := strings.Split(string(fileContent), "\n")
    var lineNumsStr []string
    var lineNums []int
    var safeFound = 0
    for i := 0; i < len(lines) - 1; i++ {
        lineNumsStr = strings.Split(lines[i], " ")
        /* Two ways of changing the slice type I found are using the "unsafe"
           package to break the type system (nopenopenope) or using
           "encoding/binary" to reencode it (looping seems simpler here) */
        lineNums = nil
        for j := 0; j < len(lineNumsStr); j++ {
            convertedInt, err := strconv.Atoi(lineNumsStr[j])
            if err != nil {
                log.Fatal(err)
            }
            lineNums = append(lineNums, convertedInt)
        }
        if isSafe(lineNums) || isSafeWithoutOne(lineNums) {
            safeFound++
        }
    }
    fmt.Println("Safe:", safeFound)
}
