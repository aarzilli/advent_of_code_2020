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
	
	p := point{ 0, 0}
	
	trees := 0
	
	for {
		if access(p) == '#' {
			trees++
		}
		p.i++
		p.j += 3
		if p.i >= len(M) {
			break
		}
	}
	
	fmt.Printf("trees %d\n", trees)
}
