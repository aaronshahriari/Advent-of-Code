package parts

import "fmt"

// STEPS:
// 1. compare char (input) by char (map) ???

// NAME EVERYTHING HERE STARTING WITH CAP
func NumberReader(file_string string) string{
    mapped_nums := map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    test_slice := []string{
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen",
    }


    for _, i := range test_slice {
        for _, j := range i {
            for r, l := range mapped_nums {
                for _, t := range r {
                    fmt.Println("test_slice letter:", string(j))
                    fmt.Println("mapped_num letter:", string(t))
                    fmt.Println("mapped_num number:", string(l))
                    if j == t {
                        fmt.Println("MATCHED")
                        continue
                    } else { fmt.Println("BREAK"); break }
                }
                break
            }
            break
        }
        break
    }
    return ""
    // return file_string
}
