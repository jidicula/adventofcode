package main

import (
	"adventofcode"
	"fmt"
	"strconv"
)

func main() {
	inputs := adventofcode.ReadInput("input.txt")
	// inputs := adventofcode.ReadInput("test.txt")

	elf := 0
	max := 0
	for _, v := range inputs {
		if v == "" {
			if elf > max {
				max = elf
			}

			elf = 0
		}
		cal, _ := strconv.Atoi(v)
		elf += cal
	}

	fmt.Printf("max elf cal: %d\n", max)

	part2(inputs)
}

// part2 ...
func part2(inputs []string) {
	elf := 0
	top3 := []int{0, 0, 0}
	for _, v := range inputs {
		if v == "" {
			if elf > top3[1] {
				if elf > top3[2] {
					top3[0] = top3[1]
					top3[1] = top3[2]
					top3[2] = elf
				} else {
					top3[0] = top3[1]
					top3[1] = elf
				}
			} else {
				if elf > top3[0] {
					top3[0] = elf
				}
			}

			elf = 0
		}
		cal, _ := strconv.Atoi(v)
		elf += cal
	}

	fmt.Printf("sum of top 3: %d\n", top3[0]+top3[1]+top3[2])

}
