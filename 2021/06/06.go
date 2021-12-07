package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode"
)

type FishTimers [9]uint64

// increment adds 1 day to the state and returns the number of new fish.
func (ft *FishTimers) increment() uint64 {
	tmp := ft[0]

	for i := 0; i <= 8; i++ {
		switch i {
		case 8:
			ft[8] = tmp
		case 6:
			ft[6] = ft[7] + tmp
		default:
			ft[i] = ft[i+1]
		}
	}
	return tmp
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")
	initialFishS := strings.Split(inputs[0], ",")
	var fishStates FishTimers
	// initial load
	var count uint64
	for _, s := range initialFishS {
		t, err := strconv.Atoi(s)
		if err != nil {
			panic("Couldn't convert, RIP")
		}
		fishStates[t] += 1
		count++
	}

	// part 1
	for i := 1; i <= 80; i++ {
		count += fishStates.increment()
	}

	fmt.Printf("count: %d\n", count)
	fmt.Printf("part 2: %d\n", part2(initialFishS))
}

func part2(initFishS []string) uint64 {
	var fishStates FishTimers
	// initial load
	var count uint64
	for _, s := range initFishS {
		t, err := strconv.Atoi(s)
		if err != nil {
			panic("Couldn't convert, RIP")
		}
		fishStates[t] += 1
		count++
	}

	for i := 1; i <= 256; i++ {
		count += fishStates.increment()
	}
	return count
}
