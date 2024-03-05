package parts

import (
    "log"
    "strconv"
    "unicode"
)

// CONSTRAINTS:
// 1. no digit? -> ???
// 2. one digit? -> **USE THAT DIGIT TWICE**
// 3. multiple digits? -> **FIRST AND LAST**

// MUST BE CAPS
func IsDigit(slice_output []string) int {
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

// MUST BE CAPS
func Opt_isDigit(content_str string) int {
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
