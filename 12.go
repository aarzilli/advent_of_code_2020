package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// convert string to integer
func atoi(in string) int {
	n, err := strconv.Atoi(in)
	must(err)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

type point struct {
	i, j int
}


// PART 1 STATE
var (
	dir = 0
	// 0 = East
	// 1 = North
	// 2 = West
	// 3 = South
	
	pos = point{ 0, 0 }
)

func part1(order byte, n int) {
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
		dir = ((dir-n)+4)%4
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
		default:
			panic("blah")
		}
		goto processOrder
	}
}

// PART 2 STATE
var (
	spos = point{ 0, 0 }
	wpos = point{ -1, 10 }
)

func part2(order byte, n int) {
	switch order {
	case 'N':
		wpos.i -= n
	case 'S':
		wpos.i += n
	case 'E':
		wpos.j += n
	case 'W':
		wpos.j -= n
	case 'L':
		n = n / 90
		for i := 0; i < n; i++ {
			wpos.i, wpos.j = -wpos.j, wpos.i
		}
	case 'R':
		n = n / 90
		for i := 0; i < n; i++ {
			wpos.i, wpos.j = wpos.j, -wpos.i
		}
	case 'F':
		spos.i += n*wpos.i
		spos.j += n*wpos.j
	}
}

func main() {
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
		
		part1(order, n)
		part2(order, n)
		
	}
	
	pf("PART 1: %d\n", abs(pos.i)+abs(pos.j))
	pf("PART 2: %d\n", abs(spos.i)+abs(spos.j))
}
