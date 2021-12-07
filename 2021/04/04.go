package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"adventofcode"
)

type card struct {
	values      map[int]cardValue
	row         []int
	col         []int
	unmarkedSum int
}

// getRowCol checks if a number is in a card.
func (c card) getRowCol(n int) (cardValue, bool) {
	cv, ok := c.values[n]
	return cv, ok
}

// stamp stamps a number on the bingo card.
func (c *card) stamp(row int, col int, n int) {
	c.unmarkedSum -= n
	// update row
	c.row[row]++
	// update col
	c.col[col]++
}

// checkBingo checks if card has bingo.
func (c card) checkBingo() (int, bool) {
	for i := 0; i < 5; i++ {
		if c.row[i] == 5 || c.col[i] == 5 {
			return c.unmarkedSum, true
		}
	}
	return c.unmarkedSum, false
}

type cardValue struct {
	row int
	col int
}

func main() {
	// inputs := adventofcode.ReadInput("test.txt")
	inputs := adventofcode.ReadInput("input2.txt")
	calls, cards := processInput(inputs)

	partTwo(calls, cards)

	// fmt.Printf("draws %v\n", calls)
	// fmt.Printf("cards %v\n", cards)

	// bingo game
	for _, call := range calls {
		for i := range cards {
			// check if value is in a card
			cv, ok := cards[i].getRowCol(call)
			if ok {
				cards[i].stamp(cv.row, cv.col, call)
			}
			unmarkedSum, ok := cards[i].checkBingo()

			if ok {
				// fmt.Printf("card %d\n", i)

				fmt.Printf("sum: %d\ncall: %d\n %d\n", unmarkedSum, call, unmarkedSum*call)
				os.Exit(0)
			}
		}
	}
}

func processInput(inputs []string) ([]int, []card) {
	// number calls
	calls := strings.Split(inputs[0], ",")
	var callsInt []int
	for _, s := range calls {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't convert number draw %s", s)
			os.Exit(1)
		}
		callsInt = append(callsInt, n)
	}

	// bingo cards
	var cards []card
	r := regexp.MustCompile(`(?m)\s+`)

	for line := 2; line < len(inputs); line += 6 {
		c := card{
			values:      map[int]cardValue{},
			row:         make([]int, 5),
			col:         make([]int, 5),
			unmarkedSum: 0,
		}
		for row := 0; row < 5; row++ {
			rowString := inputs[line+row]
			rowString = strings.TrimLeft(rowString, " ")
			rowString = string(r.ReplaceAll([]byte(rowString), []byte(" ")))
			rowVals := strings.Split(rowString, " ")
			if len(rowVals) > 5 {
				fmt.Printf("%v\n", rowVals)
			}
			// fmt.Printf("%v\n", rowVals)

			for col := 0; col < len(rowVals); col++ {
				val := rowVals[col]
				if val == "" {
					continue
				}
				// fmt.Printf("'%s' ", val)

				n, err := strconv.Atoi(val)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Couldn't convert bingo card \n\trow %d, col %d: %s\n", row, col, val)
					os.Exit(1)
				}
				c.values[n] = cardValue{
					row: row,
					col: col,
				}
				c.unmarkedSum += n
			}
		}
		// fmt.Printf("card start score: %d\n", c.score)

		cards = append(cards, c)
		// fmt.Printf("\n")

	}

	return callsInt, cards
}

func partTwo(calls []int, cards []card) {
	for _, call := range calls {
		var nonWinningCards []card
		for i := range cards {
			// check if value is in a card
			cv, ok := cards[i].getRowCol(call)
			if ok {
				cards[i].stamp(cv.row, cv.col, call)
			}
			unmarkedSum, ok := cards[i].checkBingo()

			if ok {
				if len(cards) == 1 {
					fmt.Printf("last winner score: %d\n", unmarkedSum*call)
				}
			} else {
				nonWinningCards = append(nonWinningCards, cards[i])
			}
		}
		cards = nonWinningCards
	}
}
