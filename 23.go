package main

import (
	"fmt"
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

//var input = 389125467 // example
var input = 467528193 // real input

type Node struct {
	n int
	next *Node
}

var state *Node
var max int
var findnode = map[int]*Node{}
var curnode *Node

func printstate() {
	cur := state
	for {
		if cur == curnode {
			pf("(%d) ", cur.n)
		} else {
			pf("%d ", cur.n)
		}
		cur = cur.next
		if cur == state {
			break
		}
	}
	pf("\n")
}

func move() {
	rem := curnode.next
	curnode.next = curnode.next.next.next.next
	tgt := curnode.n
	for {
		tgt--
		if tgt < 1 {
			tgt = max
		}
		if !(tgt == rem.n || tgt == rem.next.n || tgt == rem.next.next.n) {
			break
		}
	}
	
	if debug {
		pf("pickup %d %d %d\n", rem.n, rem.next.n, rem.next.next.n)
		pf("target %d\n", tgt)
	}
	tgtnode := findnode[tgt]
	oldnext := tgtnode.next
	tgtnode.next = rem
	rem.next.next.next = oldnext
	curnode = curnode.next
}

func printfromlabel1() {
	start := findnode[1]
	cur := start.next
	for {
		pf("%d", cur.n)
		cur = cur.next
		if cur == start {
			break
		}
	}
	pf("\n")
}

func count() int {
	cur := state
	tot := 0
	for {
		tot++
		cur = cur.next
		if cur == state {
			break
		}
	}
	return tot
}

const debug = false
const part2 = true

func main() {
	v := vatoi(splitandclean(fmt.Sprintf("%d", input), "", -1))
	lastnext := &state
	for i := range v {
		n := &Node{n: v[i]}
		findnode[v[i]] = n
		*lastnext = n
		lastnext = &n.next
		if v[i] > max || max == 0 {
			max = v[i]
		}
	}
	
	if part2 {
		total := 1000000
		for i := max+1; i < total+1; i++ {
			n := &Node{n: i}
			findnode[i] = n
			*lastnext = n
			lastnext = &n.next
			max = i
		}
	}
	
	
	*lastnext = state
	curnode = state
	
	pf("number: %d\n", count())
	
	rounds := 100
	if part2 {
		rounds = 10000000
	}
	
	for round := 0; round < rounds; round++ {
		if round % 1000 == 0 {
			pf("progress %g\n", (float64(round)/float64(rounds)) * 100)
		}
		if debug {
			pf("-- move %d --\n", round+1)
			printstate()
		}
		move()
		if debug {
			pf("\n")
		}
	}
	
	if !part2 {
		pf("final:\n")
		printstate()
		
		printfromlabel1()
	} else {
		pf("%d %d = %d\n", findnode[1].next.n, findnode[1].next.next.n, findnode[1].next.n * findnode[1].next.next.n)
	}
	
}
