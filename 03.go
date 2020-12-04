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

var M [][]byte

type point struct {
	i, j int
}

func access(p point) byte {
	row := M[p.i%len(M)]
	return row[p.j%len(row)]
}

func count(inc point) int {
	trees := 0
	p := point{0, 0}
	for {
		if access(p) == '#' {
			trees++
		}
		p.i += inc.i
		p.j += inc.j
		if p.i >= len(M) {
			break
		}
	}
	return trees
}

func main() {
	buf, err := ioutil.ReadFile("03.txt")
	must(err)
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		M = append(M, []byte(line))
	}

	p1 := count(point{1, 3})

	fmt.Printf("PART1: %d\n", p1)

	a := count(point{1, 1})
	b := count(point{1, 3})
	c := count(point{1, 5})
	d := count(point{1, 7})
	e := count(point{2, 1})

	fmt.Printf("PART2: %d\n", a*b*c*d*e)
}
