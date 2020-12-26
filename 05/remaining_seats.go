package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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
			continue
		}

		if err != nil {
			fmt.Printf("lines from file %s counted: %d\n", fn, line_count)
			return buffer[:line_count], nil
		}
	}
}

func str_to_num(line string, ones byte) float64 {
	line_len := len(line) - 1
	var number float64 = 0
	for ind, char := range line {
		power := line_len - ind
		if byte(char) == ones {
			number = number + math.Pow(2, float64(power))
		}
	}
	return number
}

func get_seat_id(line string) int {
	row_str, column_str := line[:7], line[7:]
	row, column := str_to_num(row_str, 'B'), str_to_num(column_str, 'R')
	return int(row)*8 + int(column)
}

func main() {

	file_name := "input.txt"
	lines, err := lineReader(file_name)
	if err != nil {
		fmt.Printf("Error reading %s: %v", file_name, err)
	}
	line_num := len(lines)

	seat_id := make([]int, line_num)
	for ind, line := range lines {
		seat_id[ind] = get_seat_id(line)
	}

	sort.Ints(seat_id)

	for ind, seat := range seat_id {
		if ind >= (line_num)-1 {
			break
		}

		if seat+1 != seat_id[ind+1] {
			fmt.Printf("Your seat: %d \n", seat+1)
			break
		}
	}
}
