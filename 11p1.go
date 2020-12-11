package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

var M [][]byte
var M1 [][]byte

func getseat(i, j int) byte {
	if i < 0 || i >= len(M) {
		return '.'
	}
	if j < 0 || j >= len(M[i]) {
		return '.'
	}
	return M[i][j]
}

func getocc(i, j int) int {
	if getseat(i, j) == '#' {
		return 1
	}
	return 0
}

func occupied(i, j int) int {
	return getocc(i-1, j-1) + getocc(i-1, j) + getocc(i-1, j+1) + getocc(i, j-1) + getocc(i, j+1) + getocc(i+1, j-1) + getocc(i+1, j) + getocc(i+1, j+1)
}

func step1(i, j int) {
	if M[i][j] == 'L' && occupied(i, j) == 0 {
		M1[i][j] = '#'
	} else if M[i][j] == '#' && occupied(i, j) >= 4 {
		M1[i][j] = 'L'
	} else {
		M1[i][j] = M[i][j]
	}
}

func step() {
	for i := range M {
		for j := range M[i] {
			step1(i, j)
		}
	}
	M, M1 = M1, M
}

func nochange() bool {
	for i := range M {
		for j := range M[i] {
			if M[i][j] != M1[i][j] {
				return false
			}
		}
	}
	return true
}

func allocc() int {
	count := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

const debug = false

func main() {
	buf, err := ioutil.ReadFile("11.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		M = append(M, []byte(line))
	}

	M1 = make([][]byte, len(M))
	for i := range M {
		M1[i] = make([]byte, len(M[i]))
	}

	for n := 0; n < 100000; n++ {
		pf("%d %d\n", n, allocc())
		step()
		if debug {
			for i := range M {
				fmt.Printf("%s\n", string(M[i]))
			}
		}
		if nochange() {
			break
		}
	}

	pf("PART 1: %d\n", allocc())

}
