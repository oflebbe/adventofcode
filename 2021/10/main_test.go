package main

import "testing"

func TestScore(t *testing.T) {

	testInput := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	sc := parse(testInput)
	if sc != 26397 {
		t.Errorf("got %d", sc)
	}

	if 288957 != parse2(testInput) {
		t.Errorf("got2 %d", sc)
	}
}
