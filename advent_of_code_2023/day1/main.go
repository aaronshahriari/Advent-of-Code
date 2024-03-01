package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
    "strings"
)

// CONSTRAINTS:
// 1. no digit? -> ???
// 2. one digit? -> **USE THAT DIGIT TWICE**
// 3. multiple digits? -> **FIRST AND LAST**

// Aaron's Logic
func readtxt(file_path string) []byte {
    // read txt in to content (string) variable
    content, err := os.ReadFile(file_path)
    if err != nil {
        log.Fatal(err)
    }

    return content
}

func isDigit(slice_output []string) int {
    total := 0
    for _, line := range slice_output {
        if len(line) == 0 {
            continue
        }
        digit_slice := []string{}
        for _, rune := range line {
            if unicode.IsDigit(rune) == true {
                digit_slice = append(digit_slice, string(rune))
            }
        }
        concat_digit := digit_slice[0] + digit_slice[len(digit_slice)-1]
        digit_total, err := strconv.Atoi(concat_digit)
        if err != nil {
            log.Fatal(err)
        }
        total += digit_total
    }
    return total
}

func opt_isDigit(content_str string) int {
    total := 0
    digit_slice := []string{}
    for _, rune := range content_str {
        if unicode.IsDigit(rune) == true {
            digit_slice = append(digit_slice, string(rune))
        } else if unicode.IsSpace(rune) {
            // combine first and last runes
            concat_digit := digit_slice[0] + digit_slice[len(digit_slice)-1]
            // add to running total
            digit_total, err := strconv.Atoi(concat_digit)
            if err != nil {
                log.Fatal(err)
            }
            // clear the current slice
            digit_slice = nil
            total += digit_total
        }
    }
    return total
}

func main() {
    // starting time to check optimization
    start := time.Now()

    // read in txt; convert to string; create slice
    file_path := "input.txt"

    // GLOBAL - read function (most efficient read method)
    content := readtxt(file_path)

    // AARON'S VERSION
    // get string output and convert to slice
    slice_output := strings.Split(string(content), "\n")
    total := isDigit(slice_output)

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
