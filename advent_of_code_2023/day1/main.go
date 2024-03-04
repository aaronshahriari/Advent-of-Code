package main

import (
	"fmt"
	"time"
    "strings"
    "day1/parts"
)

func part1() {
    // PART 1
    // starting time to check optimization
    start := time.Now()

    // read in txt; convert to string; create slice
    file_path := "input.txt"

    // GLOBAL - read function (most efficient read method)
    content := parts.Readtxt(file_path)

    // AARON'S VERSION
    // get string output and convert to slice
    slice_output := strings.Split(string(content), "\n")
    total := parts.IsDigit(slice_output)

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

func part2() {
    // starting time to check optimization
    start := time.Now()

    // read in txt; convert to string; create slice
    file_path := "input.txt"

    // GLOBAL - read function (most efficient read method)
    content := parts.Readtxt(file_path)

    // get string output and convert to slice
    // slice_output := strings.Split(string(content), "\n")
    // output := parts.NumberReader(slice_output)

    letters_to_nums := parts.NumberReader(string(content))

    slice_output := strings.Split(letters_to_nums, "\n")
    total := parts.IsDigit(slice_output)

    // GLOBAL - print total
    fmt.Println()
    fmt.Println("TOTAL:", total)

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}

func main() {
    // part1()
    part2()
}
