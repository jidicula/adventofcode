package adventofcode

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInput reads each line of a text file into a string slice.
func ReadInput(filepath string) []string {
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
