package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

var enumcache = map[int]int{}

func enum(input []int, last int) int {
	if len(input) == 1 {
		return 1
	}
	if r, ok := enumcache[len(input)]; ok {
		return r
	}
	n := 0
	for i := range input {
		if input[i]-last <= 3 {
			n += enum(input[i+1:], input[i])
		} else if input[i]-last > 3 {
			break
		}
	}
	enumcache[len(input)] = n
	return n
}

var input []int

func main() {
	buf, err := ioutil.ReadFile("10.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		input = append(input, atoi(line))
	}

	sort.Ints(input)

	jolt := map[int]int{}
	last := 0
	for i := range input {
		d := input[i] - last
		last = input[i]

		switch d {
		case 1, 2, 3:
			jolt[d]++
		default:
			panic("blah")
		}
	}
	jolt[3]++

	pf("PART 1: %d %d %d\n", jolt[1], jolt[3], jolt[1]*jolt[3])

	count := enum(append(input, input[len(input)-1]+3), 0)
	pf("PART 2: %d\n", count)
}
