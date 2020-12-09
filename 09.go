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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

var input []int

func valid(n int, start, end int) bool {
	for i := start; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if input[i]+input[j] == n {
				return true
			}
		}
	}
	return false
}

func validate(preamble int) int {
	for i := preamble; i < len(input); i++ {
		if !valid(input[i], i-preamble, i) {
			pf("PART 1: %d\n", input[i])
			return input[i]
		}
	}
	panic("bad")
}

func search(tgt int) {
	for start := 0; start < len(input); start++ {
		for radius := 2; start+radius < len(input); radius++ {
			if sum(start, radius, tgt) == tgt {
				pf("INTERVAL: %d %d\n", start, radius)
				min, max := minmax(start, radius)
				pf("MIN %d MAX %d\n", min, max)
				pf("PART 2: %d\n", min+max)
			}
		}
	}
}

func fastersearch(tgt int) {
	prevcumsum := make([]int, len(input))
	for i := range input {
		prevcumsum[i] = input[i]
	}

	for radius := 2; radius < len(input); radius++ {
		for start := 0; start < len(input); start++ {
			if start+radius > len(input) {
				continue
			}
			s := prevcumsum[start] + input[start+radius-1]
			prevcumsum[start] = s
			if s == tgt {
				pf("INTERVAL: %d %d\n", start, radius)
				min, max := minmax(start, radius)
				pf("MIN %d MAX %d\n", min, max)
				pf("PART 2: %d\n", min+max)
				return
			}
		}
	}
}

func sum(start, radius, tgt int) int {
	if start+radius >= len(input) {
		return -1
	}
	r := 0
	for i := 0; i < radius; i++ {
		r += input[i+start]
		if r > tgt {
			return -1
		}
	}
	return r
}

func minmax(start, radius int) (min, max int) {
	min = -1
	for i := 0; i < radius; i++ {
		if min < 0 {
			min = input[i+start]
		}
		if input[i+start] < min {
			min = input[i+start]
		}
		if input[i+start] > max {
			max = input[i+start]
		}
	}
	return
}

func main() {
	buf, err := ioutil.ReadFile("09.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		input = append(input, atoi(line))
	}

	tgt := validate(25)
	//tgt := validate(5)
	//search(tgt)
	//pf("\n")
	fastersearch(tgt)
}
