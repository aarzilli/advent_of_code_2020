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

type rule struct {
	kind string
	rngs [][2]int
}

type ticket struct {
	cnt []int
	valid bool
}

var rules []rule
var yourticket []int
var tickets []*ticket

func satisfiesAny(n int) bool {
	for i := range rules {
		if satisfies(n, &rules[i]) {
			return true
		}
	}
	return false
}

func satisfies(n int, r *rule) bool {
	for i := range r.rngs {
		if n >= r.rngs[i][0] && n <= r.rngs[i][1] {
			return true
		}
	}
	return false
}

func main() {
	buf, err := ioutil.ReadFile("16.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	phase := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		switch phase {
		case 0:
			if line == "your ticket:" {
				phase = 1
				break
			}
			v := splitandclean(line, ":", -1)
			var r rule
			r.kind = v[0]
			w := splitandclean(v[1], " or ", -1)
			for _, x := range w {
				ww := vatoi(splitandclean(x, "-", -1))
				r.rngs = append(r.rngs, [2]int{ ww[0], ww[1] })
			}
			rules = append(rules, r)
		case 1:
			yourticket = vatoi(splitandclean(line, ",", -1))
			phase = 2
		
		case 2:
			if line != "nearby tickets:" {
				panic("blah")
			}
			phase = 3
		
		case 3:
			var t ticket
			t.cnt = vatoi(splitandclean(line, ",", -1))
			t.valid = true
			tickets = append(tickets, &t)
		}
	}
	
	bad := []int{}
	
	for i := range tickets {
		for j := range tickets[i].cnt {
			if !satisfiesAny(tickets[i].cnt[j]) {
				bad = append(bad, tickets[i].cnt[j])
				tickets[i].valid = false
			}
		}
	}
	
	determined := make([]string, len(tickets[0].cnt))
	
	allowed := make([]map[string]bool, len(tickets[0].cnt))
	for i := range allowed {
		allowed[i] = map[string]bool{}
		for k := range rules {
			allowed[i][rules[k].kind] = true
		}
	}
	
	for i := range tickets {
		if !tickets[i].valid {
			continue
		}
		for j := range tickets[i].cnt {
			for k := range rules {
				if !allowed[j][rules[k].kind] {
					continue
				}
				if !satisfies(tickets[i].cnt[j], &rules[k]) {
					delete(allowed[j], rules[k].kind)
				}
			}
		}
	}
	
	didsomething := true
	for didsomething {
		didsomething = false
		for j := range tickets[0].cnt {
			if determined[j] != "" {
				continue
			}
			//pf("%d %#v\n", j, allowed[j])
			if len(allowed[j]) == 1 {
				for k := range allowed[j] {
					determined[j] = k
					pf("position %d is %s\n", j, determined[j])
				}
				for k := range tickets[0].cnt {
					if k == j {
						continue
					}
					delete(allowed[k], determined[j])
				}
				didsomething = true
				break
			}
		}
	}
	
	tot := 0
	for i := range bad {
		tot += bad[i]
	}
	pf("PART 1: %d\n", tot)
	
	m := 1
	for i := range yourticket {
		if strings.HasPrefix(determined[i], "departure") {
			m *= yourticket[i]
			pf("%d\n", yourticket[i])
		}
	}
	pf("PART 2: %d\n", m)
	
}
