package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// lineReader return list of entries and its length or err
func lineReader(fn string) ([]int, int, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	lineCounter := 0
	entries := make([]int, 200)
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf(" > Failed with error: %v\n", err)
			return nil, 0, err
		}
		line = strings.TrimSuffix(line, "\n")

		// Process the line here.
		number, numbErr := strconv.Atoi(line)
		if numbErr != nil {
			fmt.Printf(" > Parsing Number failed (%s): %v", line, numbErr)
			return nil, 0, numbErr
		}
		entries[lineCounter] = number
		lineCounter++

		if err != nil {
			return entries, lineCounter, nil
		}
	}
}

func main() {

	entries, entriesNumb, err := lineReader("entries.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("entries (%d): %d \n", entriesNumb, entries)
	for st_ind, st_elem := range entries {
		for snd_ind, snd_elem := range entries[st_ind+1:] {

			// Part Two
			for _, trd_elem := range entries[snd_ind+1:] {
				if st_elem+snd_elem+trd_elem == 2020 {
					fmt.Printf("Value triplet: (%d, %d, %d)\t", st_elem, snd_elem, trd_elem)
					fmt.Printf("Product: %d\n", st_elem*snd_elem*trd_elem)
				}
			}

			// Part one
			if st_elem+snd_elem == 2020 {
				fmt.Printf("Value Pair: (%d, %d)\t", st_elem, snd_elem)
				fmt.Printf("Product: %d", st_elem*snd_elem)
			}
		}
	}
}
