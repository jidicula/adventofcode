package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type cave string

func (c cave) isSmall() bool {
	for _, r := range c {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// pathCount counts distinct paths.
func (c cave) pathCount(cm caveMap, traversed map[cave]struct{}) int {
	// fmt.Printf("%s\n", c)
	if c == "end" {
		return 1
	}
	var count int
	traversed[c] = struct{}{}
	for _, nextCave := range cm[c] {
		_, ok := traversed[nextCave]
		if ok && nextCave.isSmall() {
			continue
		}

		traversedCopy := make(map[cave]struct{})
		for k, v := range traversed {
			traversedCopy[k] = v
		}
		count += nextCave.pathCount(cm, traversedCopy)
	}
	return count
}

func (c cave) pathsWith2Visits(cm caveMap, canVisitTwice cave, alreadyVisitedTwice bool, path string, solutions map[string]struct{}) {
	if c == "end" {
		solutions[path] = struct{}{}
		return
	}

	for _, to := range cm[c] {
		secondVisit := alreadyVisitedTwice
		visited := strings.Contains(path, string(to))
		if visited && to.isSmall() {
			if to == canVisitTwice {
				if alreadyVisitedTwice {
					continue
				}
				secondVisit = true
			} else {
				// already traversed
				continue
			}
		}
		to.pathsWith2Visits(cm, canVisitTwice, secondVisit, fmt.Sprintf("%s,%s", path, c), solutions)
	}
}

type caveMap map[cave][]cave

func main() {
	inputs := strings.Split(input, "\n")
	cm := make(caveMap)
	for _, s := range inputs[:len(inputs)-1] {
		nodes := strings.Split(s, "-")
		src := cave(nodes[0])
		dest := cave(nodes[1])
		cm[src] = append(cm[src], dest)
		cm[dest] = append(cm[dest], src)
	}

	var count int
	for _, c := range cm["start"] {
		count += c.pathCount(cm, map[cave]struct{}{"start": {}})
	}
	fmt.Printf("part 1: %d\n", count)
	solutions := make(map[string]struct{})
	for n := range cm {
		if n == "start" || n == "end" || !n.isSmall() {
			continue
		}

		for _, c := range cm["start"] {
			c.pathsWith2Visits(cm, n, false, "start", solutions)
		}
	}

	fmt.Printf("part 2: %d\n", len(solutions))
}
