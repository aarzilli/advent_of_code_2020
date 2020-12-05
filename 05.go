package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

const debugconvertex = false

func convertex(m int, s, e int, sel string) (int, string) {
	if debugconvertex {
		pf("convertex %d %d %s\n", s, e-1, sel)
	}
	first, second := byte('F'), byte('B')
	if m != 0 {
		first, second = 'L', 'R'
	}
	if len(sel) == 0 {
		if s != e-1 {
			panic("blah 2")
		}
		return s, sel
	}
	switch sel[0] {
	case first:
		return convertex(m, s, s+(e-s)/2, sel[1:])
	case second:
		return convertex(m, s+(e-s)/2, e, sel[1:])
	default:
		if s != e-1 {
			panic("blah")
		}
		return s, sel
	}
}

func convert(s string) (row, col int) {
	var rest string
	row, rest = convertex(0, 0, 128, s)
	col, _ = convertex(1, 0, 8, rest)
	return row, col
}

func rowid(row, col int) int {
	return row*8 + col
}

func main() {
	maxid := 0

	found := map[int]bool{}

	buf, err := ioutil.ReadFile("05.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		row, col := convert(line)
		id := rowid(row, col)
		found[id] = true

		if id > maxid {
			maxid = id
		}
	}
	pf("PART 1: %d\n", maxid)

	for row := 0; row <= 127; row++ {
		for col := 0; col <= 7; col++ {
			id := rowid(row, col)
			if !found[id] && found[id-1] && found[id+1] {
				pf("PART 2: %d\n", id)
			}
		}
	}
}
