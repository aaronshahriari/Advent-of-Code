package parts

import (
	"strings"
)

// STEPS:
// 1. compare char (input) by char (map) ???

// NAME EVERYTHING HERE STARTING WITH CAP
func NumberReader(file_string string) string{
    mapped_nums := map[string]string{
        "one": "one1one",
        "two": "two2two",
        "three": "three3three",
        "four": "four4four",
        "five": "five5five",
        "six": "six6six",
        "seven": "seven7seven",
        "eight": "eight8eight",
        "nine": "nine9nine",
    }

    for old, new := range mapped_nums {
        file_string = strings.Replace(file_string, old, new, -1)
    }

    return file_string
}
