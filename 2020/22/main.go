package main

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func round(player1 []byte, player2 []byte) (r1 []byte, r2 []byte) {
	a := player1[0]
	b := player2[0]
	r1 = player1[1:]
	r2 = player2[1:]
	if a > b {
		r1 = append(r1, a)
		r1 = append(r1, b)
	} else {
		r2 = append(r2, b)
		r2 = append(r2, a)
	}
	return
}

func roundRecurse(player1 []byte, player2 []byte, memo map[string]struct{}) (r1 []byte, r2 []byte) {

	/*for _, p := range player1 {
		fmt.Printf("%d,", p)
	}
	fmt.Println()
	for _, p := range player2 {
		fmt.Printf("%d,", p)
	}
	fmt.Println()
	fmt.Println()*/
	a := player1[0]
	b := player2[0]
	r1 = player1[1:]
	r2 = player2[1:]

	if int(a) > len(r1) || int(b) > len(r2) {
		if a > b {
			r1 = append(r1, a)
			r1 = append(r1, b)
		} else {
			r2 = append(r2, b)
			r2 = append(r2, a)
		}
		return
	}

	// clone slice trick from so
	s1 := append([]byte{}, r1[0:a]...)
	s2 := append([]byte{}, r2[0:b]...)
	_, p1Wins := recurseGame(s1, s2)

	r1 = player1[1:]
	r2 = player2[1:]
	if p1Wins {
		r1 = append(r1, a)
		r1 = append(r1, b)
	} else {
		r2 = append(r2, b)
		r2 = append(r2, a)
	}
	return
}

func readPlayer(lines string) (res []byte) {
	linesSlice := strings.Split(lines, "\n")

	for i, line := range linesSlice {
		if i == 0 {
			if !strings.HasPrefix(line, "Player ") {
				panic("Parse error")
			}
			continue
		}
		if line == "" {
			continue
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic("read error")
			}
			res = append(res, byte(num))
		}
	}
	return
}

func readInput(input string) ([]byte, []byte) {
	var players [2][]byte
	for i, block := range strings.Split(input, "\n\n") {
		players[i] = readPlayer(block)
	}
	if len(players[0]) != len(players[1]) {
		panic("inconsistency")
	}
	return players[0], players[1]
}

func scoref(player []byte) (res int) {

	for i := 0; i < len(player); i++ {
		res += int(player[i]) * (len(player) - i)
	}
	return
}

func part1() int {
	fh, _ := os.Open("input.txt")
	defer fh.Close()
	buf, _ := ioutil.ReadAll(fh)
	p1, p2 := readInput(string(buf))
	for len(p1) > 0 && len(p2) > 0 {
		p1, p2 = round(p1, p2)
	}
	var s int
	if len(p1) > 0 {
		s = scoref(p1)
	} else {
		s = scoref(p2)
	}
	return s
}

func recurseGame(p1, p2 []byte) (score int, p1Wins bool) {

	memo := make(map[string]struct{})

	for len(p1) > 0 && len(p2) > 0 {
		h := md5.New()
		h.Write(p1)
		zero := []byte{0}
		h.Write(zero)
		h.Write(p2)

		st1 := string(h.Sum(nil))

		_, found := memo[st1]

		if found {
			p1Wins = true
			return
		}
		memo[st1] = struct{}{}
		p1, p2 = roundRecurse(p1, p2, memo)
	}

	p1Wins = len(p1) > 0
	if p1Wins {
		score = scoref(p1)
	} else {
		score = scoref(p2)
	}
	return
}

func main() {
	println(part1())

	fh, _ := os.Open("input.txt")
	defer fh.Close()
	buf, _ := ioutil.ReadAll(fh)
	p1, p2 := readInput(string(buf))
	score, _ := recurseGame(p1, p2)
	println(score)
}
