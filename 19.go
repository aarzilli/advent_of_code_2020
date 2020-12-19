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

type Rule struct {
	id int
	Subrules [][]int
	Match byte
}

var Rules = map[int]*Rule{}

func (r *Rule) match(line string) (ok bool, rest string) {
	if r.Match != 0 {
		if len(line) > 0 && line[0] == r.Match {
			return true, line[1:]
		}
		return false, ""
	}
	
	switch r.id {
	case 0:
		count42 := 0
		rest := line
		rests := []string{ rest }
		for {
			ok, rest = Rules[42].match(rest)
			if !ok {
				break
			}
			rests = append(rests, rest)
			count42++
		}
		
		pf("\trule 42 can be matched %d times\n", count42)
		for i := range rests {
			pf("\t\t%d %d\n", i, len(rests[i]))
		}
		
		for len(rests) > 0 {
			pf("\ttrying %d %q (%d)\n", len(rests), rests[len(rests)-1], len(rests[len(rests)-1]))
			rest := rests[len(rests)-1]
			rests = rests[:len(rests)-1]
			
			count31 := 0
			for len(rest) > 0 {
				ok, rest = Rules[31].match(rest)
				if !ok {
					count31 = -1
					break
				}
				count31++
			}
			
			pf("\trule 31 can be matched %d times (rest = %d)\n", count31, len(rest))
			if count31 > 0 && rest == "" && count42-count31 > 0 {
				pf("\tret true with %d+%d\n", count42, count31)
				return true, ""
			}
		}
		
		return false, ""
	
	default:
		for i := range r.Subrules {
			rest := line
			subok := true
			for _, id := range r.Subrules[i] {
				var subsubok bool
				subsubok, rest = Rules[id].match(rest)
				if !subsubok {
					subok = false
					break
				}
			}
			
			if subok {
				return true, rest
			}
		}
	
		return false, ""
	}
}

func main() {
	buf, err := ioutil.ReadFile("19p2.txt")
	must(err)
	blocks := strings.Split(string(buf), "\n\n")
	for _, line := range splitandclean(blocks[0], "\n", -1) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := splitandclean(line, ":", -1)
		id := atoi(v[0])
		
		rule := &Rule{ id: id }
		
		if v[1][0] == '"' {
			rule.Match = v[1][1]
		} else {
			w := splitandclean(v[1], "|", -1)
			for _, field := range w {
				seq := vatoi(splitandclean(field, " ", -1))
				rule.Subrules = append(rule.Subrules, seq)
			}
		}
		
		pf("%#v\n", rule)
		
		Rules[rule.id] = rule
	}
	
	n := 0
	for _, line := range splitandclean(blocks[1], "\n", -1) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		pf("matching %q\n", line)
		if ok, rest := Rules[0].match(line); ok && rest == "" {
			pf("\tmatch\n")
			n++
		}
	}
	
	pf("PART 1: %d\n", n)
}
