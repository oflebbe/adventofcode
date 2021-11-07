package main

import "fmt"

func Valid(chars []rune) bool {
	lastChar := ' '
	num := 0
	doubleChar := []rune{}
	straightFound := false
	for _, c := range chars {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
		if lastChar == c {
			doubleChar = append(doubleChar, c)
		}
		lastChar++

		if c == lastChar {
			num++
			if num == 2 {
				straightFound = true
			}
		} else {
			lastChar = c
			num = 0
		}
	}
	if !straightFound {
		return false
	}
	if len(doubleChar) != 2 {
		return false
	}

	return doubleChar[0] != doubleChar[1]
}

func Next(chars []rune) []rune {
first:
	j := len(chars) - 1
	for j >= 0 {
		switch chars[j] {
		case 'i' - 1, 'l' - 1, 'o' - 1:
			chars[j] += 2
			if Valid(chars) {
				return chars
			}
			goto first
		case 'z':
			chars[j] = 'a'
			j--
		default:
			chars[j]++
			if Valid(chars) {
				return chars
			}
			goto first
		}
	}
	return nil
}

func main() {
	fmt.Printf("%s\n", string(Next([]rune("hxbxwxba"))))
	fmt.Printf("%s\n", string(Next(Next([]rune("hxbxwxba")))))
}
