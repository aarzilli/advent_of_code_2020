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

var tiles = []*Tile{}

type Tile struct {
	id int
	M  [][]byte
	n  int
}

func borders(tile *Tile) []string {
	r := []string{}
	r = append(r, string(tile.M[0]))
	r = append(r, string(tile.M[len(tile.M)-1]))

	left := []byte{}
	right := []byte{}
	for i := range tile.M {
		left = append(left, tile.M[i][0])
		right = append(right, tile.M[i][len(tile.M[i])-1])
	}

	r = append(r, string(left))
	r = append(r, string(right))
	return r
}

func reverse(s string) string {
	r := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return string(r)
}

var assoc = map[string][]conf{}

type conf struct {
	borderi int
	flip    bool
	id      int
}

func main() {
	buf, err := ioutil.ReadFile("20.txt")
	must(err)
	blocks := strings.Split(string(buf), "\n\n")
	for _, block := range blocks {
		cur := &Tile{}

		for _, line := range splitandclean(block, "\n", -1) {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if strings.HasPrefix(line, "Tile ") {
				cur.id = atoi(nolast(splitandclean(line, " ", -1)[1]))
			} else {
				cur.M = append(cur.M, []byte(line))
			}
		}

		tiles = append(tiles, cur)
	}

	for _, tile := range tiles {
		for i, border := range borders(tile) {
			assoc[border] = append(assoc[border], conf{i, false, tile.id})
			revb := reverse(border)
			assoc[revb] = append(assoc[revb], conf{i, true, tile.id})
		}
	}

	//pf("%#v\n", assoc)

	m := map[int]int{}

	for _, v := range assoc {
		if len(v) > 1 {
			for _, c := range v {
				m[c.id]++
			}
			//pf("%#v\n", v)
		}
	}

	tot := 1
	for id := range m {
		if m[id] == 4 {
			pf("%d\n", id)
			tot *= id
		}
	}

	pf("PART 1: %d\n", tot)
}
