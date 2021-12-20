package main

import (
	"adventofcode"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

// newPoint creates a new point from an "x,y" line of input.
func newPoint(inputLine string) point {
	coords := strings.Split(inputLine, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("Couldn't convert x")
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("Couldn't convert y")
	}

	return point{
		x: x,
		y: y,
	}
}

type page map[point]struct{}

func (p page) fold(ins instruction) {
	for pt := range p {
		newPt := point{
			x: pt.x,
			y: pt.y,
		}
		if ins.axis == "x" {
			if pt.x > ins.value {
				delete(p, pt)
				newPt.x -= 2 * (newPt.x - ins.value)
				p[newPt] = struct{}{}
			}
		} else if ins.axis == "y" {
			if pt.y > ins.value {
				delete(p, pt)
				newPt.y -= 2 * (newPt.y - ins.value)
				p[newPt] = struct{}{}
			}
		}
	}
}

type instruction struct {
	axis  string
	value int
}

func newInstruction(inputLine string) instruction {
	fold := strings.Split(strings.TrimLeft(inputLine, "fold along "), "=")
	val, err := strconv.Atoi(fold[1])
	if err != nil {
		panic("Couldn't convert instruction value")
	}
	return instruction{
		axis:  fold[0],
		value: val,
	}
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")
	pg := make(page)
	var inputLineNum int
	// parse coords
	for inputs[inputLineNum] != "" {
		p := newPoint(inputs[inputLineNum])
		pg[p] = struct{}{}
		inputLineNum++
	}
	// parse instructions
	var instructions []instruction
	for _, line := range inputs[inputLineNum+1:] {
		instructions = append(instructions, newInstruction(line))
	}

	// Part 1
	pg.fold(instructions[0])
	fmt.Printf("Part 1: %d\n", len(pg))

	// Part 2
	for _, ins := range instructions[1:] {
		pg.fold(ins)
	}

	for y := 0; y <= 6; y++ {
		for x := 0; x <= 40; x++ {
			p := point{
				x: x,
				y: y,
			}
			_, ok := pg[p]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
