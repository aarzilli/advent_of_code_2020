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

func verify(rng []int, ch byte, pwd string) bool {
	m := make(map[byte]int)
	for i := range pwd {
		m[pwd[i]]++
	}
	return m[ch] >= rng[0] && m[ch] <= rng[1]
}

func verify2(rng []int, ch byte, pwd string) bool {
	n := 0
	if pwd[rng[0]-1] == ch {
		n++
	}
	if pwd[rng[1]-1] == ch {
		n++
	}
	return n == 1
}

func main() {
	fmt.Printf("hello\n")
	buf, err := ioutil.ReadFile("02.txt")
	must(err)
	cnt := 0
	cnt2 := 0
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		fields := splitandclean(line, " ", -1)
		rng := vatoi(splitandclean(fields[0], "-", -1))
		ch := fields[1][0]
		pwd := fields[2]
		
		if verify(rng, ch, pwd) {
			cnt++
		}
		
		if verify2(rng, ch, pwd) {
			//fmt.Printf("valid\n")
			cnt2++
		} else {
			//fmt.Printf("invalid\n")
		}
		
	}
	fmt.Printf("PART 1: %d\n", cnt)
	fmt.Printf("PART 2: %d\n", cnt2)
}
