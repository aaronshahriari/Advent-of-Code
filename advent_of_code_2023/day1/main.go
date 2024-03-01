package main

import (
	"fmt"
	"time"
    "strings"
    "day1/part1"
)

func main() {
    // starting time to check optimization
    start := time.Now()

    // read in txt; convert to string; create slice
    file_path := "input.txt"

    // GLOBAL - read function (most efficient read method)
    content := part1.Readtxt(file_path)

    // AARON'S VERSION
    // get string output and convert to slice
    slice_output := strings.Split(string(content), "\n")
    total := part1.IsDigit(slice_output)

    // WORSE
    // OPTIMIZED VERSION
    // transpile the long string
    // total := opt_isDigit(string(content))

    // GLOBAL - print total
    fmt.Println(total)

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}
