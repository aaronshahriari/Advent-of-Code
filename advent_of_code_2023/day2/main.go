package main

import (
	"day2/lib"
	"day2/parts"
	"fmt"
    "time"
)

func part1() {
    ///// MAIN /////
    content := lib.ReadTxt("input.txt")
    fmt.Println(string(content))

    parts.Tester()
    ///// MAIN /////
}

func main() {
    // starting time to check optimization
    start := time.Now()

    part1()

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}
