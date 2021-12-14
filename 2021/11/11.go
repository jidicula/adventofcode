package main

import (
	"adventofcode"
	"fmt"
)

const gridSize = 9

type octopus struct {
	row int
	col int
}

// increment increases the energy level of an octopus.
func (g grid) increment(o octopus) {
	if g[o] != 9 {
		g[o]++
	} else {
		g[o] = 0
	}
}

type grid map[octopus]int

func (g grid) String() string {
	var s string

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			value := g[octopus{
				row: i,
				col: j,
			}]
			if value == 0 {
				s += fmt.Sprintf("\033[1m%d\033[0m", value)
			} else {
				s += fmt.Sprintf("%d", value)
			}
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

// adjacent returns a slice of octopuses that are adjacent to the provided
// coordinates.
func (g grid) adjacent(i, j int) []octopus {
	var neighbours []octopus
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	row_limit := gridSize
	if row_limit > 0 {
		column_limit := gridSize
		for x := max(0, i-1); x <= min(i+1, row_limit); x++ {
			for y := max(0, j-1); y <= min(j+1, column_limit); y++ {
				if x != i || y != j {
					neighbours = append(neighbours, octopus{
						row: x,
						col: y,
					})
				}
			}
		}
	}
	return neighbours
}

func (g grid) step() int {
	f := make(flashSet)
	for o := range g {
		f.step(o, g)
	}
	return len(f)
}

// allFlashed checks if all octopi are flashing.
func (g grid) allFlashed() bool {
	for _, v := range g {
		if v != 0 {
			return false
		}
	}
	return true
}

type flashSet map[octopus]struct{}

func (f flashSet) step(o octopus, g grid) {
	// Return if already flashed
	if _, ok := f[o]; ok {
		return
	}
	// Otherwise increment
	g.increment(o)
	// If flashed, add to flashSet
	if g[o] == 0 {
		f[o] = struct{}{}
		// Recurse into octopus's neighbours
		neighbours := g.adjacent(o.row, o.col)

		for _, octoNeighbour := range neighbours {
			f.step(octoNeighbour, g)
		}
	}
}

func main() {

	// inputs := adventofcode.ReadInput("small_test.txt")
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")

	g := make(grid)
	for i, s := range inputs {
		for j, b := range s {
			v := int(b - '0')
			o := octopus{row: i, col: j}
			g[o] = v
		}
	}
	fmt.Println(g)

	var count int
	for i := 1; i <= 100; i++ {
		count += g.step()
		fmt.Printf("after step %d:\n", i)
		fmt.Println(g)
	}
	fmt.Printf("count: %d\n", count)
	fmt.Printf("part 2: %d\n", part2(inputs))
}

func part2(inputs []string) int {
	g := make(grid)
	for i, s := range inputs {
		for j, b := range s {
			v := int(b - '0')
			o := octopus{row: i, col: j}
			g[o] = v
		}
	}

	var count int
	for !g.allFlashed() {
		g.step()
		fmt.Printf("after step %d:\n", count)
		fmt.Println(g)
		count++
	}
	return count
}
