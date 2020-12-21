package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// lineReader return list of entries and its length or err
func lineReader(fn string) ([]string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	buffer := make([]string, 1000)
	line_count := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf(" > Failed with error: %v\n", err)
			return nil, err
		}
		line = strings.TrimSuffix(line, "\n")

		// Process the line here.
		if line != "" {
			buffer[line_count] = line
			line_count++
		}

		if err != nil {
			fmt.Printf("lines from file counted: %d", line_count)
			return buffer[:line_count], nil
		}
	}
}

func main() {

	file_name := "input.txt"
	topography, err := lineReader(file_name)
	if err != nil {
		fmt.Printf("Error reading %s: %v", file_name, err)
	}

	line_length := len(topography[0])
	slopes := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	results := make([]int, len(slopes))
	for ind, slope := range slopes {
		pos, tree_hits := 0, 0
		for line_ind, tree_line := range topography {
			if line_ind%slope[0] != 0 {
				continue
			}
			if tree_line[pos] == byte('#') {
				tree_hits++
			}
			pos = (pos + slope[1]) % line_length
		}
		results[ind] = tree_hits
	}
	total := 1
	for _, val := range results {
		total = total * val
	}
	fmt.Printf("hits: %d | total: %d", results, total)
}
