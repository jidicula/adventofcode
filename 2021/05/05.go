package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode"
)

type point struct {
	x int
	y int
}

// newPoint creates a point from a string slice of x and y coordinates.
func newPoint(coords []string) point {
	if len(coords) != 2 {
		panic("More than 2 coords for point")
	}
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}

	return point{
		x: x,
		y: y,
	}
}

func (p point) String() string {
	return fmt.Sprintf("[x: %d, y: %d]", p.x, p.y)
}

type line struct {
	p1 point
	p2 point
}

// newLine creates a new line.
func newLine(sPoints []string) line {
	p1 := newPoint(strings.Split(sPoints[0], ","))
	p2 := newPoint(strings.Split(sPoints[1], ","))
	l := line{p1: p1, p2: p2}
	if l.notDiagonal() {
		if p1.x < p2.x || p1.y < p2.y {
			return line{
				p1: p1,
				p2: p2,
			}
		}
		return line{
			p1: p2,
			p2: p1,
		}
	} else {
		if p1.x < p2.x {
			return line{
				p1: p1,
				p2: p2,
			}
		}
		return line{
			p1: p2,
			p2: p1,
		}
	}
}

func (l line) String() string {
	return fmt.Sprintf("line{%s %s}", l.p1, l.p2)
}

// notDiagonal checks if a line is not diagonal (i.e. horizontal or vertical).
func (l line) notDiagonal() bool {
	return l.p1.x == l.p2.x || l.p1.y == l.p2.y
}

// is45 checks if a line is diagonal at exactly 45°
func (l line) is45() bool {
	// absDiff returns absolute difference
	absDiff := func(x, y int) int {
		if x < y {
			return y - x
		}
		return x - y
	}
	return absDiff(l.p2.x, l.p1.x) == absDiff(l.p2.y, l.p1.y)
}

// draw returns a plane with the line drawn on it.
func (l line) draw(p plane) plane {
	return p
}

type plane map[point]int

func main() {
	inputs := adventofcode.ReadInput("test.txt")
	// inputs := adventofcode.ReadInput("input.txt")

	pl := make(plane)
	for _, s := range inputs {
		s := strings.Replace(s, " ", "", -1)
		sPoints := strings.Split(s, "->")
		l := newLine(sPoints)
		if l.notDiagonal() {
			// fmt.Printf("%v\n", l)

			if l.p1.x == l.p2.x {
				for i := l.p1.y; i <= l.p2.y; i++ {
					p := point{
						x: l.p1.x,
						y: i,
					}
					pl[p]++
				}
			} else {
				for i := l.p1.x; i <= l.p2.x; i++ {
					p := point{
						x: i,
						y: l.p1.y,
					}
					pl[p]++
				}
			}
		}
	}
	// fmt.Printf("%v\n", pl)

	var count uint16
	for _, v := range pl {
		if v >= 2 {
			count++
		}
	}
	fmt.Printf("2 or more overlaps: %d\n", count)

	// part 2
	fmt.Printf("part 2 overlaps: %d\n", part2(inputs))
}

func part2(inputs []string) uint16 {
	pl := make(plane)
	for _, s := range inputs {
		s := strings.Replace(s, " ", "", -1)
		sPoints := strings.Split(s, "->")
		l := newLine(sPoints)
		if l.notDiagonal() {
			// fmt.Printf("%v\n", l)
			if l.p1.x == l.p2.x {
				for i := l.p1.y; i <= l.p2.y; i++ {
					p := point{
						x: l.p1.x,
						y: i,
					}
					pl[p]++
				}
			} else {
				for i := l.p1.x; i <= l.p2.x; i++ {
					p := point{
						x: i,
						y: l.p1.y,
					}
					pl[p]++
				}
			}

		}
		if l.is45() {
			for i := 0; i <= l.p2.x-l.p1.x; i++ {
				var y int
				if l.p2.y > l.p1.y {
					y = l.p1.y + i
				} else {
					y = l.p1.y - i
				}
				p := point{
					x: l.p1.x + i,
					y: y,
				}
				pl[p]++
			}

			// fmt.Printf("%s\n", l)
		}
	}
	// fmt.Printf("%v\n", pl)

	var count uint16
	for _, v := range pl {
		if v >= 2 {
			count++
		}
	}
	return count
}
