package main

import (
	"fmt"
	"time"
    "strings"
    "day1/parts"
    "day1/lib"
)

func part1(content []byte) int{
    // AARON'S VERSION
    // get string output and convert to slice
    slice_output := strings.Split(string(content), "\n")
    total := parts.IsDigit(slice_output)

    // WORSE
    // transpile the long string
    // total := opt_isDigit(string(content))

    return total
}

func part2(content []byte) int{
    letters_to_nums := parts.NumberReader(string(content))

    slice_output := strings.Split(letters_to_nums, "\n")
    total := parts.IsDigit(slice_output)

    return total
}

func main() {
    // starting time to check optimization
    start := time.Now()

    ///// MAIN /////
    content := lib.ReadTxt("input.txt")
    
    fmt.Println("PART 1:")
    total_p1 := part1(content)
    fmt.Println("P1 TOTAL:", total_p1)

    fmt.Println()

    fmt.Println("PART 2:")
    total_p2 := part2(content)
    fmt.Println("P2 TOTAL:", total_p2)
    ///// MAIN /////

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}
