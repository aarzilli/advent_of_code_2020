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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func transf(subjectNumber, loopSize int) int {
	value := 1
	for cnt := 0; cnt < loopSize; cnt++ {
		value *= subjectNumber
		value = value % 20201227
	}
	return value
}

func findLoopSize(subj, tgt0, tgt1 int) (loopsz, which int) {
	value := 1
	for cnt := 0; ; cnt++ {
		if value == tgt0 {
			return cnt, 0
		}
		if value == tgt1 {
			return cnt, 1
		}
		value *= subj
		value = value % 20201227
	}
}

func main() {
	buf, err := ioutil.ReadFile("25.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	input := vatoi(lines[:len(lines)-1])
	if len(input) != 2 {
		panic("blah")
	}
	pf("%v\n", input)

	loopsz, idx := findLoopSize(7, input[0], input[1])
	pf("loop size: %d (input line %d)\n", loopsz, idx)
	pf("encryption key: %d (19414467)\n", transf(input[(idx+1)%2], loopsz))
}
