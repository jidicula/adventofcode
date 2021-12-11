package main

import (
	"adventofcode"
	"fmt"
	"sort"
)

type height uint

type heightMap [][]height

// adjacent returns heights that are adjacent to the provided coordinates.
func (hm heightMap) adjacent(row, col int) []point {
	var neighbours []point

	if row > 0 {
		neighbours = append(neighbours, point{h: hm[row-1][col], row: row - 1, col: col})
	}
	if row < len(hm)-1 {
		neighbours = append(neighbours, point{h: hm[row+1][col], row: row + 1, col: col})
	}

	if col > 0 {
		neighbours = append(neighbours, point{h: hm[row][col-1], row: row, col: col - 1})
	}
	if col < len(hm[row])-1 {
		neighbours = append(neighbours, point{h: hm[row][col+1], row: row, col: col + 1})
	}

	return neighbours
}

// isLow checks if a provided coordinate is a low point.
func (hm heightMap) isLow(row, col int) bool {
	h := hm[row][col]
	adjacent := hm.adjacent(row, col)
	for _, p := range adjacent {
		if h >= p.h {
			return false
		}
	}
	return true
}

type point struct {
	h   height
	row int
	col int
}

type basin map[point]struct{}

func (b basin) size() uint {
	return uint(len(b))
}

func (b basin) floodFill(p point, hm heightMap) {
	if _, ok := b[p]; ok || p.h == 9 {
		return
	}
	b[p] = struct{}{}
	neighbours := hm.adjacent(p.row, p.col)
	// fmt.Printf("neighbours of point %d: %d\n", p.h, len(neighbours))
	// return

	for _, pn := range neighbours {
		b.floodFill(pn, hm)
	}
	return
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")

	var heightmap heightMap
	for _, s := range inputs {
		var row []height
		for _, r := range s {
			row = append(row, height(r-'0'))
		}
		heightmap = append(heightmap, row)
	}

	var riskSum uint
	var lowPoints []point
	for row := 0; row < len(heightmap); row++ {
		for col := 0; col < len(heightmap[0]); col++ {
			if heightmap.isLow(row, col) {
				p := point{h: heightmap[row][col], row: row, col: col}
				lowPoints = append(lowPoints, p)
				riskSum += uint(p.h + 1)
			}
		}
	}
	fmt.Printf("part 1: %d\n", riskSum)
	result := part2(lowPoints, heightmap)
	fmt.Printf("largest 3: %v\n", result)

}

func part2(lowPoints []point, heightmap heightMap) uint {
	var sizes []uint
	for _, p := range lowPoints {
		b := make(basin)
		b.floodFill(p, heightmap)
		sizes = append(sizes, b.size())
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	largest3 := sizes[:3]
	result := uint(1)
	for _, x := range largest3 {
		result *= x
	}
	return uint(result)
}
