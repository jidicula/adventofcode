package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode"
)

type crabSwarm map[position]int
type position int

// cost computes the total cost of moving the crabSwarm to that horizontal
// position.
func (cs crabSwarm) cost(horPos position) int {
	absDiff := func(x, y position) int {
		if x < y {
			return int(y - x)
		}
		return int(x - y)
	}
	var total int
	for p, count := range cs {
		total += count * absDiff(horPos, p)
	}
	return total
}

// costPart2 computes the total cost, using the new cost definition from Part
// 2.
func (cs crabSwarm) costPart2(horPos position) int {
	absDiff := func(x, y position) int {
		if x < y {
			return int(y - x)
		}
		return int(x - y)
	}
	singleCost := func(current, final position) int {
		distance := absDiff(current, final)
		return (distance * int(1+distance) / 2)
	}
	var total int
	for p, count := range cs {
		total += count * singleCost(p, horPos)
	}
	return total
}

// mode returns the mode of the crabSwarm
func (cs crabSwarm) mode() position {
	var mode position
	for k, v := range cs {
		if v > cs[mode] {
			mode = k
		}
	}
	return mode
}

// max returns the max of the crabSwarm.
func (cs crabSwarm) max() position {
	var max position
	for k := range cs {
		if k > max {
			max = k
		}
	}
	return max
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")

	inputIntS := strings.Split(inputs[0], ",")

	positions := make(crabSwarm)
	for _, s := range inputIntS {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic("Couldn't convert RIP")
		}
		p := position(x)
		positions[p]++
	}

	// start with mode
	minPos := positions.mode()
	minCost := positions.cost(minPos)

	// determine descent direction
	newPos := minPos + 1
	needIncrease := positions.cost(newPos) < minCost

	newCost := minCost
	var steps uint
	for newCost <= minCost {
		steps++
		minCost = newCost
		minPos = newPos
		if needIncrease {
			newPos++
		} else {
			newPos--
		}
		newCost = positions.cost(newPos)
	}
	fmt.Printf("min cost: %d, new position: %d, in %d steps\n", minCost, minPos, steps)
	part2minCost, part2minPos, part2steps := part2(positions)
	fmt.Printf("min cost: %d, new position: %d, in %d steps\n", part2minCost, part2minPos, part2steps)
}

func part2(positions crabSwarm) (int, position, uint) {
	// start with mode
	minPos := positions.mode()
	minCost := positions.costPart2(minPos)

	// determine descent direction
	newPos := minPos + 1
	needIncrease := positions.costPart2(newPos) < minCost
	newCost := minCost
	var steps uint
	for newCost <= minCost {
		steps++
		minCost = newCost
		minPos = newPos
		if needIncrease {
			newPos++
		} else {
			newPos--
		}
		newCost = positions.costPart2(newPos)
	}

	return minCost, minPos, steps
}
