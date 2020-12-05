package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// returns x without the last character
func nolast(x string) string {
	return x[:len(x)-1]
}

// splits a string, trims spaces on every element
func splitandclean(in, sep string, n int) []string {
	v := strings.SplitN(in, sep, n)
	for i := range v {
		v[i] = strings.TrimSpace(v[i])
	}
	return v
}

// convert string to integer
func atoi(in string) int {
	n, err := strconv.Atoi(in)
	must(err)
	return n
}

// convert vector of strings to integer
func vatoi(in []string) []int {
	r := make([]int, len(in))
	for i := range in {
		var err error
		r[i], err = strconv.Atoi(in[i])
		must(err)
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func exit(n int) {
	os.Exit(n)
}

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

const debugconvertex = false

func convertex(m int, s, e int, sel string) (int, string) {
	if debugconvertex {
		pf("convertex %d %d %s\n", s, e-1, sel)
	}
	if m == 0 {
		switch sel[0] {
		case 'F':
			return convertex(m, s, s + (e-s)/2, sel[1:])
		case 'B':
			return convertex(m, s + (e-s)/2, e, sel[1:])
		default:
			if s != e-1 {
				panic("blah")
			}
			return s, sel
		}
	} else {
		if len(sel) == 0 {
			if s != e-1 {
				panic("blah 2")
			}
			return s, sel
		}
		switch sel[0] {
		case 'L':
			return convertex(m, s, s + (e-s)/2, sel[1:])
		case 'R':
			return convertex(m, s + (e-s)/2, e, sel[1:])
		default:
			panic("blah 3")
		}
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
	row, col := convert("BBFFBBFRLL")
	pf("%d %d %d\n", row, col, rowid(row, col))
	
	maxid := 0

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
		if id > maxid {
			maxid = id
		}
	}
	pf("PART 1: %d\n", maxid)
}
