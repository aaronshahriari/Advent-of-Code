package parts

import (
	"fmt"
	"strings"
    "strconv"
    "log"
)

// CONSTRAINTS:
// > 12 RED
// > 13 GREEN
// > 14 BLUE
// RESULT: ADD GAME NUMBER UP
func SumColors(content []string) int{
    id_sum := 0

    red_constr := 12
    green_constr := 13
    blue_constr := 14

    for game := range content {
        game_message := fmt.Sprintf("Game %v: ", game + 1)
        content[game] = strings.TrimPrefix(content[game], game_message)
        content[game] = strings.ReplaceAll(content[game], ",", "")
        sep_content := strings.Split(content[game], ";")

        for sep := range sep_content {
            single_it_map := make(map[string]string)
            parts := strings.Fields(sep_content[sep])

            // iterate over the slices
            // starting from index 1 and incrementing by 2
            for i := 1; i < len(parts); i += 2 {
                single_it_map[parts[i]] = parts[i-1]
            }

            // gets the value of the map
            // returns 0 if it does not exist
            getValue := func(key string) int{
                if value, ok := single_it_map[key]; ok {
                    value, err := strconv.Atoi(value)
                    if err != nil {
                        log.Fatal(err)
                    }
                    return value
                }
                return 0
            }

            fmt.Println(single_it_map)

            // check if current map is not 0 then continue
            if len(single_it_map) == 0 {
            } else {
                red_count := getValue("red")
                green_count := getValue("green")
                blue_count := getValue("blue")
                if (red_count < red_constr &&
                green_count < green_constr &&
                blue_count < blue_constr) {
                    fmt.Println("CHECKS OUT")
                    continue
                } else { fmt.Println("DOESN'T CHECKS OUT") }
            }
        }
    }
    return id_sum
}
