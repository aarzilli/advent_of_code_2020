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

type point struct {
	i, j int
}

// 0 = East
// 1 = North
// 2 = West
// 3 = South

func main() {
	dir := 0
	pos := point{ 0, 0 }
	buf, err := ioutil.ReadFile("12.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		order := line[0]
		n := atoi(line[1:])
		pf("%c %d\n", order, n)
		
		processOrder:
		switch order {
		case 'N':
			pos.i -= n
		case 'S':
			pos.i += n
		case 'E':
			pos.j += n
		case 'W':
			pos.j -= n
		case 'L':
			n = n / 90
			dir = (dir+n)%4
		case 'R':
			n = n / 90
			dir -= n
			switch dir {
			case -1:
				dir = 3
			case -2:
				dir = 2
			case -3:
				dir = 1
			}
		case 'F':
			switch dir {
			case 0: // east
				order = 'E'
			case 1: // north
				order = 'N'
			case 2: // west
				order = 'W'
			case 3: // south
				order = 'S'
			}
			goto processOrder
		}
	}
	
	pf("%d %d\n", pos.i, pos.j)
	pf("PART 1: %d\n", abs(pos.i)+abs(pos.j))
}
