package main

import (
	"fmt"
	"io/ioutil"
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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func main() {
	buf, err := ioutil.ReadFile("06.txt")
	must(err)
	lines := strings.Split(string(buf), "\n\n")
	tot := 0
	tot2 := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		m := map[byte]int{}
		people := splitandclean(line, "\n", -1)
		for _, subln := range people {
			for _, ch := range subln {
				m[byte(ch)]++
			}
		}

		for _, v := range m {
			if v == len(people) {
				tot2++
			}
		}
		tot += len(m)
	}

	pf("PART 1: %d\n", tot)
	pf("PART 2: %d\n", tot2)
}
