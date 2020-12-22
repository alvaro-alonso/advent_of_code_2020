package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr, iyr, eyr      int
	hgt, hcl, ecl, pid string
}

func parse_line(line string, pass *Passport) {
	pairs := strings.Split(line, " ")
	for _, pair := range pairs {
		key_val := strings.Split(pair, ":")
		key, val := key_val[0], key_val[1]
		switch key {
		case "byr":
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("failed parsing password with input: %s", line)
			}
			pass.byr = num
		case "iyr":
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("failed parsing password with input: %s", line)
			}
			pass.iyr = num
		case "eyr":
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("failed parsing password with input: %s", line)
			}
			pass.eyr = num
		case "hgt":
			pass.hgt = val
		case "hcl":
			pass.hcl = val
		case "ecl":
			pass.ecl = val
		case "pid":
			pass.pid = val
		}
	}
}

func valid_pass(pass Passport) bool {
	if pass.byr < 1920 || pass.byr > 2002 {
		fmt.Printf("Wrong year: %d\n", pass.byr)
		return false
	}
	if pass.iyr < 2010 || pass.iyr > 2020 {
		return false
	}
	if pass.eyr < 2020 || pass.eyr > 2030 {
		return false
	}

	if len(pass.hgt) < 2 {
		return false
	}
	hgt_len := len(pass.hgt) - 2
	hgt_units := pass.hgt[hgt_len:]
	hgt_num, _ := strconv.Atoi(pass.hgt[:hgt_len])
	if hgt_units != "cm" && hgt_units != "in" {
		return false
	} else if hgt_units == "cm" && (hgt_num < 150 || hgt_num > 193) {
		return false
	} else if hgt_units == "in" && (hgt_num < 59 || hgt_num > 76) {
		return false
	}

	if pass.hcl == "" {
		return false
	} else if pass.hcl[0] != byte('#') || len(pass.hcl) != 7 {
		return false
	}

	_, pid_err := strconv.Atoi(pass.pid)
	if len(pass.pid) != 9 || pid_err != nil {
		fmt.Println("number is to big")
		return false
	}

	switch pass.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		fmt.Println("false pass.ecl")
		return false
	}
}

// lineReader return list of entries and its length or err
func lineReader(fn string) ([]Passport, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	buffer := make([]Passport, 1000)
	passport_count := 0
	passport := new(Passport)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf(" > Failed with error: %v\n", err)
			return nil, err
		}
		line = strings.TrimSuffix(line, "\n")

		// Process the line here.
		if line != "" {
			parse_line(line, passport)
			continue
		}

		buffer[passport_count] = *passport
		passport = new(Passport)
		passport_count++

		if err != nil {
			fmt.Printf("lines from file %s counted: %d\n", fn, passport_count)
			return buffer[:passport_count], nil
		}
	}
}

func main() {

	file_name := "input.txt"
	passports, err := lineReader(file_name)
	if err != nil {
		fmt.Printf("Error reading %s: %v", file_name, err)
	}

	valid_passports := 0
	for _, passport := range passports {
		if valid_pass(passport) == true {
			valid_passports++
		}
	}
	fmt.Printf("Number of valid passports: %d\n", valid_passports)
}
