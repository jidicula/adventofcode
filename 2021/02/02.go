package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"adventofcode"
)

type Instruction struct {
	direction string
	delta     int
}

type Position struct {
	horizontal int
	depth      int
}

// move changes position.
func (p *Position) move(ins Instruction) {
	switch ins.direction {
	case "down":
		p.depth += ins.delta
	case "up":
		p.depth -= ins.delta
	case "forward":
		p.horizontal += ins.delta
	}
}

func main() {
	inputs := adventofcode.ReadInput("input.txt")
	// inputs := adventofcode.ReadInput("test.txt")
	var pos Position
	for _, s := range inputs {
		split := strings.Split(s, " ")
		if len(split) < 2 {
			continue
		}

		d, err := strconv.Atoi(split[1])
		if err != nil {
			os.Exit(1)
		}
		instruct := Instruction{
			direction: split[0],
			delta:     d,
		}
		pos.move(instruct)
	}
	fmt.Printf("%v\n", pos.horizontal*pos.depth)
	partTwo(inputs)
}

type PositionPt2 struct {
	Position
	aim int
}

// move changes PositionPt2
func (p *PositionPt2) move(ins Instruction) {
	switch ins.direction {
	case "down":
		p.aim += ins.delta
	case "up":
		p.aim -= ins.delta
	case "forward":
		p.horizontal += ins.delta
		p.depth += p.aim * ins.delta
	}
}

func partTwo(inputs []string) {
	var pos PositionPt2
	for _, s := range inputs {
		split := strings.Split(s, " ")
		if len(split) < 2 {
			continue
		}

		d, err := strconv.Atoi(split[1])
		if err != nil {
			os.Exit(1)
		}
		instruct := Instruction{
			direction: split[0],
			delta:     d,
		}
		pos.move(instruct)
	}
	fmt.Printf("%v\n", pos.horizontal*pos.depth)
}
