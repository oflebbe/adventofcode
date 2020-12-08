package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func valid(input string, what string) bool {
	switch what {
	case "byr":
		v, _ := strconv.Atoi(input)
		return v >= 1920 && v <= 2002

	case "iyr":
		v, _ := strconv.Atoi(input)
		return v >= 2010 && v <= 2020
	case "eyr":
		v, _ := strconv.Atoi(input)
		return v >= 2020 && v <= 2030

	case "pid":
		b, _ := regexp.Match(`^\d{9}$`, []byte(input))
		return b
	case "hcl":
		b, _ := regexp.Match(`^#[0-9a-f]{6}$`, []byte(input))
		return b

	case "ecl":
		b, _ := regexp.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, []byte(input))
		return b

	case "hgt":
		v, err := strconv.Atoi(input[0 : len(input)-2])
		if err != nil {
			return false
		}
		if strings.HasSuffix(input, "cm") {
			return v >= 150 && v <= 193
		} else if strings.HasSuffix(input, "in") {
			return v >= 59 && v <= 76
		}
		return false
	case "cid":
		return true
	default:
		return false
	}

}

func Marshal(input string) error {
	count := 0
	for _, tok := range strings.Fields(input) {
		keyVals := strings.SplitN(tok, ":", 2)

		field := keyVals[0]
		if field == "cid" {
			continue
		}
		if !valid(keyVals[1], field) {
			return errors.New("validation fail")
		}

		count++

	}
	if count == 7 {
		return nil
	}
	return errors.New("missing smthg")

}

func process(fn string) int {
	fh, _ := os.Open(fn)
	defer fh.Close()
	buf, _ := ioutil.ReadAll(fh)

	count := 0
	for _, lines := range strings.Split(string(buf), "\n\n") {

		err := Marshal(lines)
		if err == nil {
			count++
		}

	}

	return count
}

func main() {
	input := `pid:839955293
byr:1928 hcl:#fffffd ecl:hzl iyr:2011
hgt:162cm eyr:2023`

	e := Marshal(input)
	if e != nil {
		log.Fatal(e)
	}

	if process("wrong.txt") != 0 {
		log.Fatal("wrong wrong")
	}

	if process("right.txt") != 4 {
		log.Fatal("right wrong")
	}

	fmt.Println(process("input.txt"))

}
