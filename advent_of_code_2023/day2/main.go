package main

import (
	"day2/lib"
	"day2/parts"
	"fmt"
	"strings"
	"time"
)

func part1(content []string) int{
    total := parts.SumColors(content)
    return total
}

func main() {
    // starting time to check optimization
    start := time.Now()

    // content := lib.ReadTxt("inputs/sample_p1.txt")
    content := lib.ReadTxt("inputs/input.txt")
    split_content := strings.Split(string(content), "\n")

    ///// PART 1 /////
    total := part1(split_content)
    fmt.Println("Total:", total)

    ///// PART 2 /////

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}
