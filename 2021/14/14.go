package main

import (
	"adventofcode"
	"fmt"
	"strings"
)

type polymer struct {
	pairs         map[string]int
	elementCounts map[rune]int
}

type insertionPair struct {
	pair  string
	count int
}

// pairInsert
func (p polymer) pairInsert(rules []string) {
	var insertions []insertionPair
	// delete first, then insert additions
	for _, rule := range rules {
		pair, insertion := parseRule(rule)
		if count, ok := p.pairs[pair]; ok {
			delete(p.pairs, pair)
			left := pair[:1] + string(insertion)
			right := string(insertion) + pair[1:]
			p.elementCounts[insertion] += count
			insertions = append(insertions, insertionPair{
				pair:  left,
				count: count,
			})
			insertions = append(insertions, insertionPair{
				pair:  right,
				count: count,
			})
		}
	}
	for _, x := range insertions {
		p.pairs[x.pair] += x.count
	}
}

func (p polymer) max() (rune, int) {
	var max int
	var maxElement rune
	for k, v := range p.elementCounts {
		if v > max {
			max = v
			maxElement = k
		}
	}
	return maxElement, max
}

func (p polymer) min() (rune, int) {
	_, min := p.max()
	var minElement rune
	for k, v := range p.elementCounts {
		if v < min {
			min = v
			minElement = k
		}
	}
	return minElement, min
}

func parseRule(rule string) (string, rune) {
	tokens := strings.Split(rule, " -> ")
	return tokens[0], rune(tokens[1][0])
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")

	template := inputs[0]
	p := polymer{
		pairs:         make(map[string]int),
		elementCounts: make(map[rune]int),
	}
	// load initial template
	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		p.pairs[pair]++

	}
	for _, c := range template {
		p.elementCounts[c]++
	}

	// load rules
	rules := inputs[2:]
	// Part 1
	for i := 0; i < 10; i++ {
		p.pairInsert(rules)
	}
	_, min := p.min()
	_, max := p.max()
	fmt.Printf("Part 1: %d\n", max-min)

	// Part 2
	for i := 10; i < 40; i++ {
		p.pairInsert(rules)
	}
	_, min = p.min()
	_, max = p.max()
	fmt.Printf("Part 1: %d\n", max-min)
}
