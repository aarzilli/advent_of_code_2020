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

var input []int

func valid(n int, start, end int) bool {
	for i :=  start; i < end; i++ {
		for j := i+1; j < end; j++ {
			if input[i] + input[j] == n {
				return true
			}
		}
	}
	return false
}

func validate(preamble int) {
	for i := preamble; i < len(input); i++ {
		if !valid(input[i], i-preamble, i) {
			pf("PART 1: %d\n", input[i])
			return
		}
	}
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
	
	validate(25)
}
