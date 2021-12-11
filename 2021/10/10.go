package main

import (
	"adventofcode"
	"container/list"
	"fmt"
	"sort"
)

type chunk string

// isValid returns the last checked byte in a chunk and whether chunk is valid
// or not.
func (c chunk) isValid() (*list.List, byte, bool) {
	stack := list.New()
	var last byte
	for _, b := range c {
		last = byte(b)
		// if left paren
		if _, ok := pairs[last]; ok {
			stack.PushBack(last)
		} else {
			if pop, ok := stack.Remove(stack.Back()).(byte); ok {
				if last != pairs[pop] {
					return stack, last, false
				}
			}
		}
	}
	return stack, last, true
}

var points = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var pairs = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var part2points = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")
	var score int
	var incomplete []*list.List
	for _, s := range inputs {
		remaining, b, ok := chunk(s).isValid()
		if !ok {
			score += points[b]
		} else {
			incomplete = append(incomplete, remaining)
		}
	}
	fmt.Printf("part 1: %d\n", score)
	var part2Scores []int
	for _, l := range incomplete {
		part2Scores = append(part2Scores, getScore(l))
	}
	sort.Slice(part2Scores, func(i, j int) bool { return part2Scores[i] < part2Scores[j] })

	fmt.Printf("part 2: %d\n", part2Scores[len(part2Scores)/2])
}

func getScore(l *list.List) int {
	var score int
	for l.Back() != nil {
		if pop, ok := l.Remove(l.Back()).(byte); ok {
			score = score*5 + part2points[pairs[pop]]
		}
	}
	return score
}
