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

func convert(line string) int {
	r := 0
	for i := range line {
		r <<= 1
		if line[i] == 'B' || line[i] == 'R' {
			r++
		}
	}
	return r
}

func main() {
	maxid := 0
	found := map[int]bool{}

	buf, err := ioutil.ReadFile("05.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		id := convert(line)
		found[id] = true

		if id > maxid {
			maxid = id
		}
	}

	fmt.Printf("PART 1: %d\n", maxid)

	for id := 0; id <= (127*8 + 7); id++ {
		if !found[id] && found[id-1] && found[id+1] {
			fmt.Printf("PART 2: %d\n", id)
			break
		}
	}
}
