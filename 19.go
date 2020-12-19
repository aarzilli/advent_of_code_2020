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

func splitandclean(in, sep string, n int) []string {
	v := strings.SplitN(in, sep, n)
	for i := range v {
		v[i] = strings.TrimSpace(v[i])
	}
	return v
}

func atoi(in string) int {
	n, err := strconv.Atoi(in)
	must(err)
	return n
}

func vatoi(in []string) []int {
	r := make([]int, len(in))
	for i := range in {
		var err error
		r[i], err = strconv.Atoi(in[i])
		must(err)
	}
	return r
}

type Rule struct {
	id       int
	Subrules [][]int
	Match    byte
}

var Rules = map[int]*Rule{}

func (r *Rule) match(line string) (ok bool, rest string) {
	if r.Match != 0 {
		if len(line) > 0 && line[0] == r.Match {
			return true, line[1:]
		}
		return false, ""
	}

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

func part2match(line string) bool {
	count42 := 0
	rest := line
	rests := []string{rest}
	for {
		var ok bool
		ok, rest = Rules[42].match(rest)
		if !ok {
			break
		}
		rests = append(rests, rest)
		count42++
	}

	for len(rests) > 0 {
		rest := rests[len(rests)-1]
		rests = rests[:len(rests)-1]

		count31 := 0
		for len(rest) > 0 {
			var ok bool
			ok, rest = Rules[31].match(rest)
			if !ok {
				count31 = -1
				break
			}
			count31++
		}

		if count31 > 0 && rest == "" && count42-count31 > 0 {
			return true
		}
	}

	return false
}

func countmatch(text string, match func(string) bool) int {
	n := 0
	for _, line := range splitandclean(text, "\n", -1) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if match(line) {
			n++
		}
	}
	return n
}

func main() {
	buf, err := ioutil.ReadFile("19.txt")
	must(err)
	blocks := strings.Split(string(buf), "\n\n")
	for _, line := range splitandclean(blocks[0], "\n", -1) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := splitandclean(line, ":", -1)
		id := atoi(v[0])

		rule := &Rule{id: id}

		if v[1][0] == '"' {
			rule.Match = v[1][1]
		} else {
			w := splitandclean(v[1], "|", -1)
			for _, field := range w {
				seq := vatoi(splitandclean(field, " ", -1))
				rule.Subrules = append(rule.Subrules, seq)
			}
		}

		Rules[rule.id] = rule
	}

	fmt.Printf("PART 1: %d\n", countmatch(blocks[1], func(line string) bool {
		ok, rest := Rules[0].match(line)
		return ok && rest == ""
	}))

	fmt.Printf("PART 2: %d\n", countmatch(blocks[1], part2match))
}
