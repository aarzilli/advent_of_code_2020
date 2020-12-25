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

func transf(subjectNumber, loopSize int) int {
	value := 1
	for cnt := 0; cnt < loopSize; cnt++ {
		value *= subjectNumber
		value = value % 20201227
	}
	return value
}

func findLoopSize(subj, tgt int) int {
	value := 1
	for cnt := 0; ; cnt++ {
		if value == tgt {
			return cnt
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
	
	/*cardpub := 5764801
	doorpub := 17807724*/
	
	cardpub := input[0]
	doorpub := input[1]
	
	cardls := findLoopSize(7, cardpub)
	pf("card loop size: %d\n", cardls)
	//pf("door loop size: %d\n", findLoopSize(7, doorpub))
	
	pf("encryption key: %d\n", transf(doorpub, cardls))
}
