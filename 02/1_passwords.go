package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	min_occ   int
	max_occ   int
	character byte
	password  string
}

func loadPolicy(args []string) (*Policy, error) {
	policy := new(Policy)
	min_occ, err := strconv.Atoi(args[0])
	max_occ, max_err := strconv.Atoi(args[1])
	if err == nil {
		err = max_err
	}

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	policy.min_occ = min_occ
	policy.max_occ = max_occ
	policy.character = args[2][0]
	policy.password = args[3]

	return policy, nil
}

func rightPolicy(pol *Policy) bool {
	count := strings.Count(pol.password, string(pol.character))
	if pol.min_occ <= count && count <= pol.max_occ {
		return true
	}
	return false
}

// lineReader return list of entries and its length or err
func lineReader(fn string) (uint, error) {
	file, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	re := regexp.MustCompile(`(\d+)|(\w+)`)

	var correct_passwords uint = 0
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf(" > Failed with error: %v\n", err)
			return 0, err
		}
		line = strings.TrimSuffix(line, "\n")
		fmt.Printf(line)

		// Process the line here.
		r := re.FindAllString(line, 4)
		fmt.Printf("regex: %s\n", r)
		policy, pol_err := loadPolicy(r)
		if pol_err != nil {
			fmt.Printf(" error parsing line: %s", line)
			return 0, err
		}
		correct_policy := rightPolicy(policy)
		if correct_policy == true {
			correct_passwords++
		}

		if err != nil {
			fmt.Printf(" > Failed with error: %v\n", err)
			return correct_passwords, nil
		}
	}
}

func main() {

	correct_passwords, err := lineReader("passwords.txt")
	if err != nil {
		fmt.Printf("Error counting correct passwords: %v", err)
	}
	fmt.Printf("Total Number of correct passwords: %d", correct_passwords)
}
