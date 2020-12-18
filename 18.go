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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

type node interface{}

type expr struct {
	v []node
}

func parse(line string) node {
	toks := splitandclean(strings.Replace(line, " ", "", -1), "", -1)
	r, toks := parseToks(toks)
	if len(toks) != 0 {
		panic(fmt.Errorf("something left to parse %q", toks))
	}
	return r
}

func parseToks(toks []string) (node, []string) {
	childs := []node{}
	for len(toks) > 0 {
		if toks[0] == "(" {
			var child node
			child, toks = parseToks(toks[1:])
			childs = append(childs, child)
			if toks[0] != ")" {
				panic("parenthesis mismatch")
			}
			toks = toks[1:]
		} else if toks[0] == ")" {
			return &expr{childs}, toks
		} else {
			childs = append(childs, toks[0])
			toks = toks[1:]
		}
	}
	return &expr{childs}, []string{}
}

func eval1(root node) int {
	switch n := root.(type) {
	case string:
		return atoi(n)
	case *expr:
		r := eval1(n.v[0])
		for i := 1; i < len(n.v); i += 2 {
			op := n.v[i].(string)
			opnd := eval1(n.v[i+1])
			switch op {
			case "+":
				r += opnd
			case "*":
				r *= opnd
			default:
				panic("blah")
			}
		}
		return r

	default:
		panic("blah")
	}

	return 0
}

func eval2One(v *[]node, kind string) bool {
	for i := 1; i < len(*v); i += 2 {
		if (*v)[i] == kind {
			newv := make([]node, 0, len(*v))

			newv = append(newv, (*v)[:i-1]...)
			x := eval2((*v)[i-1])
			y := eval2((*v)[i+1])
			r := 0
			switch (*v)[i] {
			case "+":
				r = x + y
			case "*":
				r = x * y
			default:
				panic("blah")
			}
			newv = append(newv, r)
			newv = append(newv, (*v)[i+2:]...)
			*v = newv
			return true
		}
	}
	return false
}

func eval2(root node) int {
	switch n := root.(type) {
	case string:
		return atoi(n)
	case int:
		return n
	case *expr:
		v := append([]node{}, n.v...)
		for eval2One(&v, "+") {
		}
		for eval2One(&v, "*") {
		}
		if len(v) != 1 {
			panic("blah")
		}
		return eval2(v[0])

	default:
		panic("blah")
	}

	return 0
}

func main() {
	buf, err := ioutil.ReadFile("18.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	tot1, tot2 := 0, 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		expr := parse(line)
		tot1 += eval1(expr)
		tot2 += eval2(expr)
	}
	pf("PART 1: %d\n", tot1)
	pf("PART 2: %d\n", tot2)
}
