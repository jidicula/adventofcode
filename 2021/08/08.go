package main

import (
	"adventofcode"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type digit map[byte]struct{}

// Bytes returns a slice of bytes in a digit.
func (d digit) Bytes() []byte {
	var b []byte
	for c := range d {
		b = append(b, c)
	}
	return b
}

func (d digit) String() string {
	var s string
	for c := range d {
		s += fmt.Sprintf("%c", c)
	}
	return s
}

// newDigit creates a new digit from a string
func newDigit(s string) digit {
	d := make(digit)
	for _, b := range s {
		d[byte(b)] = struct{}{}
	}
	return d
}

var digitDefinitions = []digit{
	{'a': {}, 'b': {}, 'c': {}, 'e': {}, 'f': {}, 'g': {}},
	{'c': {}, 'f': {}},
	{'a': {}, 'c': {}, 'd': {}, 'e': {}, 'g': {}},
	{'a': {}, 'c': {}, 'd': {}, 'f': {}, 'g': {}},
	{'b': {}, 'c': {}, 'd': {}, 'f': {}},
	{'a': {}, 'b': {}, 'd': {}, 'f': {}, 'g': {}},
	{'a': {}, 'b': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}},
	{'a': {}, 'c': {}, 'f': {}},
	{'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}},
	{'a': {}, 'b': {}, 'c': {}, 'd': {}, 'f': {}, 'g': {}},
}

// intersection returns a digit mask of segments present in both digits.
func (d digit) intersection(d2 digit) digit {
	both := make(digit)
	for c := range d {
		if _, ok := d2[c]; ok {
			both[c] = struct{}{}
		}
	}
	return both
}

// union returns a digit mask of segments present in either digit.
func (d digit) union(d2 digit) digit {
	either := d2
	for c := range d {
		either[c] = struct{}{}
	}
	return either
}

// notIn returns a digit mask of bytes from d not in d2.
func (d digit) notIn(d2 digit) digit {
	onlyInD := make(digit)
	for k := range d {
		if _, ok := d2[k]; !ok {
			onlyInD[k] = struct{}{}
		}
	}
	return onlyInD
}

// key is encoded segment, value is decoded segment.
type mapping map[byte]byte

func (m mapping) String() string {
	s := "mapping[\n"
	for k, v := range m {
		s += fmt.Sprintf("\t%c:%c\n", k, v)
	}
	s += "]"
	return s
}

// isComplete checks if mapping is isComplete
func (m mapping) isComplete() bool {
	for i := byte('a'); i <= byte('g'); i++ {
		if v := m[i]; v == 0 {
			return false
		}
	}
	return len(m) != 0
}

func main() {
	// inputs := adventofcode.ReadInput("single-digit.txt")
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input.txt")
	var count uint
	for _, s := range inputs {
		output := strings.Split(s, " | ")[1]
		digits := strings.Split(output, " ")
		for _, d := range digits {
			switch len(d) {
			case len(digitDefinitions[1]), len(digitDefinitions[4]), len(digitDefinitions[7]), len(digitDefinitions[8]):
				count++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", count)
	result := part2(inputs)
	fmt.Printf("Part 2: %d\n", result)
}

func part2(inputs []string) uint {
	var sum uint
	for _, l := range inputs {
		m := solveMapping(l)
		output := strings.Split(l, " | ")[1]
		outputDigits := strings.Split(output, " ")
		var decodedOutput string
		for _, od := range outputDigits {
			d := make(digit)
			for _, b := range od {
				decoded := m[byte(b)]
				d[decoded] = struct{}{}
			}

			for i, definition := range digitDefinitions {
				if reflect.DeepEqual(definition, d) {
					decodedOutput += fmt.Sprintf("%d", i)
				}
			}
		}

		v, err := strconv.Atoi(decodedOutput)
		if err != nil {
			panic("Couldn't convert RIP")
		}
		sum += uint(v)
	}
	return sum
}

func solveMapping(inputLine string) mapping {
	signals := strings.Split(inputLine, " | ")[0]

	sigLengths := make(map[int][]digit)
	for _, s := range strings.Split(signals, " ") {
		sigLengths[len(s)] = append(sigLengths[len(s)], newDigit(s))
	}

	inverseMapping := make(mapping)
	for !inverseMapping.isComplete() {
		chars := digitDefinitions[8]
		// chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
		for char := range chars {
			_, ok := inverseMapping[char]
			if ok {
				continue
			}

			switch char {
			case 'a':
				seven := sigLengths[3][0]
				one := sigLengths[2][0]
				if a := seven.notIn(one).Bytes(); len(a) == 1 {
					inverseMapping[char] = a[0]
				}
			case 'c':
				zeroSixNine := sigLengths[6]
				one := sigLengths[2][0]
				for _, d := range zeroSixNine {
					if inOne := one.notIn(d).Bytes(); len(inOne) == 1 {
						c := inOne[0]
						inverseMapping[char] = c
					}
				}
			case 'f':
				one := sigLengths[2][0]
				mask := digit{inverseMapping['c']: {}}
				if f := one.notIn(mask).Bytes(); len(f) == 1 {
					inverseMapping[char] = f[0]
				}
			case 'd':
				twoThreeFive := sigLengths[5]
				and := twoThreeFive[0]
				for _, d := range twoThreeFive[1:] {
					and = and.intersection(d)
				}
				four := sigLengths[4][0]
				one := sigLengths[2][0]
				mask := four.notIn(one)
				if d := and.intersection(mask).Bytes(); len(d) == 1 {
					inverseMapping[char] = d[0]
				}
			case 'b':
				four := sigLengths[4][0]
				one := sigLengths[2][0]
				mask := four.notIn(one)
				d := digit{inverseMapping['d']: {}}
				if b := mask.notIn(d).Bytes(); len(b) == 1 {
					inverseMapping[char] = b[0]
				}
			case 'g':
				mask := digit{inverseMapping['a']: {}}.union(
					digit{inverseMapping['b']: {}}).union(
					digit{inverseMapping['c']: {}}).union(
					digit{inverseMapping['d']: {}}).union(
					digit{inverseMapping['f']: {}},
				)

				zeroSixNine := sigLengths[6]

				for _, d := range zeroSixNine {
					inOne := d.notIn(mask).Bytes()
					if len(inOne) == 1 {
						g := inOne[0]
						inverseMapping[char] = g
					}
				}
			case 'e':
				mask := digit{inverseMapping['a']: {}}.union(
					digit{inverseMapping['b']: {}}).union(
					digit{inverseMapping['c']: {}}).union(
					digit{inverseMapping['d']: {}}).union(
					digit{inverseMapping['f']: {}}).union(
					digit{inverseMapping['g']: {}},
				)
				eight := sigLengths[7][0]
				if e := eight.notIn(mask).Bytes(); len(e) == 1 {
					inverseMapping[char] = e[0]
				}

			}
		}
	}
	m := make(mapping)
	for k, v := range inverseMapping {
		m[v] = k
	}

	return m
}
