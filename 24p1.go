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

var prefixes = []string{ "e", "se", "sw", "w", "nw", "ne" }

type Point struct {
	i, j int
}

var M = map[Point]bool{}

func main() {
	buf, err := ioutil.ReadFile("24.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		pf("line: %s\n", line)
		cur := Point{0,0}
		for line != "" {
			var instr string
			for _, pfx := range prefixes {
				if strings.HasPrefix(line, pfx) {
					instr = pfx
					line = line[len(pfx):]
					break
				}
			}
			switch instr {
			default:
				panic("blah")
			case "e":
				cur.j++
			case "se":
				cur.i--
				cur.j++
			case "sw":
				cur.i--
			case "w":
				cur.j--
			case "nw":
				cur.i++
				cur.j--
			case "ne":
				cur.i++
			}
		}
		pf("pos %v: %v -> %v\n", cur, M[cur], !M[cur])
		M[cur] = !M[cur]
	}
	
	black := 0
	for _, val := range M {
		if val {
			black++
		}
	}
	pf("PART 1: %d\n", black)
}
