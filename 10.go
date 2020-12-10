package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"sort"
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

var enumcache = map[int]int{}

func enum(input []int, last int) int {
	pf("enum(%v)\n", input)
	if len(input) == 1 {
		return 1
	}
	if r, ok := enumcache[len(input)]; ok {
		return r
	}
	n := 0
	for i := range input {
		if input[i] - last <= 3 {
			n += enum(input[i+1:], input[i])
		} else if input[i] - last > 3 {
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
	pf("%v\n", input)
	
	jolt1, jolt2, jolt3 := 0, 0, 0
	
	last := 0
	for i := range input {
		d := input[i] - last
		last = input[i]
		
		switch d {
		case 1:
			jolt1++
		case 2:
			jolt2++
		case 3:
			jolt3++
		default:
			panic("blah")
		}
	}
	
	jolt3++
	
	pf("PART 1: %d %d %d\n", jolt1, jolt3, jolt1*jolt3)
	
	count := enum(append(input, input[len(input)-1]+3), 0)
	pf("PART 2: %d\n", count)
}
