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

func enum(addr uint64, floating []int, mem map[uint64]uint64, val uint64) {
	if len(floating) == 0 {
		pf("\t%36b (decimal %d)\n", addr, addr)
		mem[addr] = val
		return
	}
	
	x := floating[0]
	bv := uint64(1<<x)
	
	enum(addr&^bv, floating[1:], mem, val)
	enum(addr|bv, floating[1:], mem, val)
}

const part2 = true

func main() {
	var maskSet uint64
	var maskClear uint64
	var floating []int
	
	mem := map[uint64]uint64{}
	
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
			pf("mask\t\t%s\n", v[1])
			
			maskClear = 0
			maskSet = 0
			
			for i := range v[1] {
				bv := uint64(1 << (len(v[1]) - i-1))
				switch v[1][i] {
				case 'X':
					// nothing
				case '1':
					maskSet |= bv
				case '0':
					maskClear |= bv
				default:
					panic("blah")
				}
			}
			
			maskFloat := ^(maskSet | maskClear) & ((1<<36)-1)
			pf("maskSet\t\t%36b\n", maskSet)
			pf("maskClear\t%36b\n", maskClear)
			
			floating = floating[:0]
			for i := 0; i < 40; i++ {
				if maskFloat & (1 << i) != 0 {
					floating = append(floating, i)
				}
			}
			
		} else {
			w := getints(v[0], false)
			addr := uint64(w[0])
			//val, _ := strconv.ParseUint(v[1], 2, 64)
			val := uint64(atoi(v[1]))
			//pf("mem[%d] = %b (%d)\n", addr, val, val)
			
			if !part2 {
				val |= maskSet
				val = val &^ maskClear
				
				//pf("\tafter mask %b\n", val)
				
				mem[addr] = val
			} else {
				pf("%36b\n", addr)
				addr |= maskSet
				//pf("%36b\n", addr)
				//addr = addr &^ maskClear
				pf("%36b (decimal %d) = (floating %v)\n", addr, addr, floating)
				enum(addr, floating, mem, val)
				
			}
		}
	}
	
		tot := uint64(0)
		for _, v := range mem {
			tot += v
		}
		pf("PART 1: %d\n", tot)
}
