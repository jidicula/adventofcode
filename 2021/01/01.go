package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput("input.txt")
	var inputInt []int

	// convert contents to int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		inputInt = append(inputInt, i)
		// fmt.Printf("%v\n", i)
	}

	var prev int
	var sum int
	for _, x := range inputInt {
		if x > prev {
			sum++
		}
		prev = x
	}
	fmt.Printf("%v\n", sum-1)

	prev = 0
	sum = 0
	for i := range inputInt[:len(inputInt)-len(inputInt)%3] {
		windowSum := inputInt[i] + inputInt[i+1] + inputInt[i+2]
		if windowSum > prev {
			sum++
		}
		prev = windowSum
	}

	fmt.Printf("%v\n", sum-1)

}

func readInput(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents
}
