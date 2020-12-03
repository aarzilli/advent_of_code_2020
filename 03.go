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

var M [][]byte

type point struct {
	i, j int
}

func access(p point) byte {
	row := M[p.i % len(M)]
	return row[p.j % len(row)]
}

func part1(inc func(*point)) int {
	trees := 0
	p := point{ 0, 0}
	for {
		if access(p) == '#' {
			trees++
		}
		inc(&p)
		if p.i >= len(M) {
			break
		}
	}
	return trees
}

func main() {
	fmt.Printf("hello\n")
	buf, err := ioutil.ReadFile("03.txt")
	must(err)
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		M = append(M, []byte(line))
	}
	
	
	p1 := part1(func(p *point) {
		p.i++
		p.j += 3
	})
	
	fmt.Printf("PART1: %d\n", p1)
	
	a := part1(func(p *point) {
		p.j++
		p.i++
	})
	
	b := part1(func(p *point) {
		p.j += 3
		p.i++
	})
	
	c := part1(func(p *point) {
		p.j += 5
		p.i++
	})
	
	d := part1(func(p *point) {
		p.j += 7
		p.i++
	})
	
	e := part1(func(p *point) {
		p.j++
		p.i += 2
	})
	
	fmt.Printf("PART2: %d\n", a*b*c*d*e)
}
