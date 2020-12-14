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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func getints(in string, hasneg bool) []int {
	v := getnums(in, hasneg, false)
	return vatoi(v)
}

func getnums(in string, hasneg, hasdot bool) []string {
	r := []string{}
	start := -1

	flush := func(end int) {
		if start < 0 {
			return
		}
		hasdigit := false
		for i := start; i < end; i++ {
			if in[i] >= '0' && in[i] <= '9' {
				hasdigit = true
				break
			}
		}
		if hasdigit {
			r = append(r, in[start:end])
		}
		start = -1
	}

	for i, ch := range in {
		isnumch := false

		switch {
		case hasneg && (ch == '-'):
			isnumch = true
		case hasdot && (ch == '.'):
			isnumch = true
		case ch >= '0' && ch <= '9':
			isnumch = true
		}

		if start >= 0 {
			if !isnumch {
				flush(i)
			}
		} else {
			if isnumch {
				start = i
			}
		}
	}
	flush(len(in))
	return r
}

var (
	maskSet   uint64
	maskClear uint64
	floating  []int
	mem1      = map[uint64]uint64{}
	mem2      = map[uint64]uint64{}
)

const debug = false

func assign1(addr, val uint64) {
	// assignment operation, part 1 version
	val |= maskSet
	val = val &^ maskClear
	mem1[addr] = val
}

func enum(addr uint64, floating []int, val uint64) {
	if len(floating) == 0 {
		if debug {
			pf("\t%36b (decimal %d)\n", addr, addr)
		}
		mem2[addr] = val
		return
	}

	bv := uint64(1 << floating[0])
	enum(addr&^bv, floating[1:], val)
	enum(addr|bv, floating[1:], val)
}

func assign2(addr, val uint64) {
	// assignment operation, part 2 version
	if debug {
		pf("%36b\n", addr)
	}
	addr |= maskSet
	if debug {
		pf("%36b (decimal %d) = (floating %v)\n", addr, addr, floating)
	}
	enum(addr, floating, val)
}

func total(mem map[uint64]uint64) uint64 {
	tot := uint64(0)
	for _, v := range mem {
		tot += v
	}
	return tot
}

func main() {
	buf, err := ioutil.ReadFile("14.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := splitandclean(line, "=", -1)
		if v[0] == "mask" {
			if debug {
				pf("mask\t\t%s\n", v[1])
			}

			maskClear = 0
			maskSet = 0
			floating = floating[:0]

			for i := range v[1] {
				bv := uint64(1 << (len(v[1]) - i - 1))
				switch v[1][i] {
				case 'X':
					floating = append(floating, (len(v[1]) - i - 1))
				case '1':
					maskSet |= bv
				case '0':
					maskClear |= bv
				default:
					panic("blah")
				}
			}

			if debug {
				pf("maskSet\t\t%36b\n", maskSet)
				pf("maskClear\t%36b\n", maskClear)
			}
		} else {
			w := getints(v[0], false)
			addr := uint64(w[0])
			val := uint64(atoi(v[1]))
			if debug {
				pf("mem[%d] = %b (%d)\n", addr, val, val)
			}

			assign1(addr, val)
			assign2(addr, val)
		}
	}

	pf("PART 1: %d\n", total(mem1))
	pf("PART 2: %d\n", total(mem2))
}
