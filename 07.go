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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func cleanup(s string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(s, ".", "", -1), "bags", "", -1), "bag", "", -1))
}

type edge struct {
	n    int
	dest string
}

var M = map[string][]edge{}

func cancontain(k string, seen map[string]bool) bool {
	seen[k] = true
	for _, e := range M[k] {
		if e.dest == "shiny gold" {
			return true
		}
		if seen[e.dest] {
			continue
		}
		if cancontain(e.dest, seen) {
			return true
		}
	}
	return false
}

func count(k string) int {
	tot := 0
	for _, s := range M[k] {
		tot += s.n * count(s.dest)
	}
	return tot + 1
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
				w[i] = cleanup(ww[1])
				M[outer] = append(M[outer], edge{dest: w[i], n: atoi(ww[0])})
			}
		} else {
			M[outer] = []edge{}
		}
	}

	n := 0
	for k := range M {
		if cancontain(k, make(map[string]bool)) {
			n++
		}
	}

	pf("PART 1: %d\n", n)
	pf("PART 2: %d\n", count("shiny gold")-1)
}
