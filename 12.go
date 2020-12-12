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

// i == y
// j == x

func main() {
	spos := point{ 0, 0 }
	wpos := point{ -1, 10 }
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
		
		//processOrder:
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
		
		pf("wpos %#v\n", wpos)
		pf("spos %#v\n", spos)
	}
	
	pf("%d %d\n", spos.i, spos.j)
	pf("PART 2: %d\n", abs(spos.i)+abs(spos.j))
	
	//pf("%d %d\n", pos.i, pos.j)
	//pf("PART 1: %d\n", abs(pos.i)+abs(pos.j))
}
