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

type node interface{}

type expr struct {
	v []node
}

func isdigit(tok string) bool {
	return tok[0] >= '0' && tok[0] <= '9'
}

func parse(line string) node {
	toks := []string{}
	v := splitandclean(line, " ", -1)
	for i := range v {
		for _, ch := range v[i] {
			toks = append(toks, fmt.Sprintf("%c", ch))
		}
	}
	pf("%q\n", toks)
	for i := range toks {
		if i+1 < len(toks) && isdigit(toks[i]) && isdigit(toks[i+1]) {
			panic("blah")
		}
	}
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
			return &expr{ childs }, toks
		} else {
			childs = append(childs, toks[0])
			toks = toks[1:]
		}
	}
	return &expr{childs}, []string{}
}

func evalOne(n *expr, kind string) bool {
	for i := 1; i < len(n.v); i += 2 {
		if n.v[i] == kind {
			pf("%#v %s %#v\n", n.v[i-1], kind, n.v[i+1])
			newv := make([]node, 0, len(n.v))
			
			newv = append(newv, n.v[:i-1]...)
			x := eval(n.v[i-1])
			y := eval(n.v[i+1])
			r := 0
			switch n.v[i] {
			case "+":
				r = x + y
			case "*":
				r = x * y
			default:
				panic("blah")
			}
			newv = append(newv, r)
			newv = append(newv, n.v[i+2:]...)
			pf("%#v -> %#v\n", n.v, newv)
			n.v = newv
			return true
		}
	}
	return false
}

func eval(root node) int {
	switch n := root.(type) {
	case string:
		return atoi(n)
	case int:
		return n
	case *expr:
		for {
			ok := evalOne(n, "+")
			if !ok {
				break
			}
		}
		for {
			ok := evalOne(n, "*")
			if !ok {
				break
			}
		}
		if len(n.v) != 1 {
			panic("blah")
		}
		return eval(n.v[0])
		/*r := eval(n.v[0])
		for i := 1; i < len(n.v); i += 2 {
			op := n.v[i].(string)
			opnd := eval(n.v[i+1])
			switch op {
			case "+":
				r += opnd
			case "*":
				r *= opnd
			default:
				panic("blah")
			}
		}
		return r*/
		
	default:
		panic("blah")
	}
	
	return 0
}

func main() {
	buf, err := ioutil.ReadFile("18.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	tot := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		expr := parse(line)
		pf("%#v\n", expr)
		x := eval(expr)
		pf("%d\n", x)
		tot += x
	}
	pf("PART 1: %d\n", tot)
}
