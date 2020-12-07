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

func cleanup(s string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(s, ".", "", -1), "bags", "", -1), "bag", "", -1))
}

var edge = map[string][]string{}
var edgecount = map[string][]int{}

func cancontain(k string, seen map[string]bool) bool {
	seen[k] = true
	for _, s := range edge[k] {
		if s == "shiny gold" {
			return true
		}
		if seen[s] {
			continue
		}
		if cancontain(s, seen) {
			return true
		}
	}
	return false
}

func count(k string) int {
	tot := 0
	for i, s := range edge[k] {
		tot += edgecount[k][i] * count(s)
	}
	pf("%s -> %d\n", k, tot)
	return tot+1
}

func main() {
	buf, err := ioutil.ReadFile("07.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		v := splitandclean(line, "contain", -1)
		outer := cleanup(v[0])
		w := splitandclean(v[1], ",", -1)
		
		if !(len(w) == 1 && w[0] == "no other bags.") {
			for i, s := range w {
				ww := splitandclean(s, " ", 2)
				_ = ww[0] // this is the number
				w[i] = cleanup(ww[1])
				
				edgecount[outer] = append(edgecount[outer], atoi(ww[0]))
				edge[outer] = append(edge[outer], w[i])
			}
		} else {
			edge[outer] = []string{}
		}
		
		//pf("%q %q\n", outer, edge[outer])
	}
	
	n := 0
	for k := range edge {
		if cancontain(k, make(map[string]bool)) {
			pf("%s\n", k)
			n++
		}
	}
	
	pf("PART 1: %d\n", n)
	
	pf("PART 2: %d\n", count("shiny gold")-1)
	
}
