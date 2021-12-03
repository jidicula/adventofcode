package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jidicula/adventofcode"
)

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")
	var intInputs [][]uint8

	for _, s := range inputs {
		var row []uint8
		for _, r := range s {
			row = append(row, uint8(r-'0'))
		}
		intInputs = append(intInputs, row)
	}

	var gamma string
	for i := 0; i < len(intInputs[0]); i++ {
		var sum int
		for j := 0; j < len(intInputs); j++ {
			sum += int(intInputs[j][i])
		}
		if sum >= len(intInputs)/2 {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}
	// fmt.Printf("gamma:\t%s\n", gamma)

	gammaInt, err := strconv.ParseUint(gamma, 2, 32)
	if err != nil {
		os.Exit(2)
	}

	var epsilon string
	for i := 0; i < len(gamma); i++ {
		if gamma[i] == byte('1') {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	// fmt.Printf("epsilon:\t%s\n", epsilon)

	epsilonInt, err := strconv.ParseUint(epsilon, 2, 12)
	if err != nil {
		os.Exit(3)
	}
	fmt.Printf("%v\n", gammaInt*epsilonInt)

	product := part2(inputs)
	fmt.Printf("%v\n", product)

}

func part2(rows []string) int {
	// oxygen
	filtered := rows
	for i := 0; i < len(rows[0]); i++ {
		colSum := 0
		mostCommon := 0
		for j := 0; j < len(filtered); j++ {
			s := fmt.Sprintf("%c", filtered[j][i])
			val, err := strconv.Atoi(s)
			if err != nil {
				os.Exit(4)
			}
			colSum += int(val)
		}
		// fmt.Printf("%v ", colSum)

		if colSum >= (len(filtered)+1)/2 {
			mostCommon = 1
		}
		filtered = filter(filtered, i, mostCommon)
		// fmt.Printf("%v: %v\n", mostCommon, filtered)
	}
	// fmt.Printf("oxygen: %v\n", filtered)

	oxygen, err := strconv.ParseUint(filtered[0], 2, 12)
	if err != nil {
		os.Exit(6)
	}

	filtered = rows
	for i := 0; len(filtered) != 1; i++ {
		colSum := 0
		leastCommon := 1
		for j := 0; j < len(filtered); j++ {
			s := fmt.Sprintf("%c", filtered[j][i])
			val, err := strconv.Atoi(s)
			if err != nil {
				os.Exit(4)
			}
			colSum += int(val)
		}
		// fmt.Printf("%v ", colSum)

		if colSum >= (len(filtered)+1)/2 {
			leastCommon = 0
		}
		filtered = filter(filtered, i, leastCommon)
		// fmt.Printf("least common %v: %v\n", leastCommon, filtered)
	}
	// fmt.Printf("co2: %v\n", filtered)

	co2, err := strconv.ParseUint(filtered[0], 2, 12)
	if err != nil {
		os.Exit(6)
	}

	return int(oxygen * co2)
}

func filter(rows []string, i int, filterBit int) []string {
	var filtered []string
	for j := 0; j < len(rows); j++ {
		s := fmt.Sprintf("%c", rows[j][i])
		v, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(5)
		}
		if v == filterBit {
			filtered = append(filtered, rows[j])
		}
	}
	return filtered
}
