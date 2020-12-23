package main

import (
	"fmt"
	"strconv"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
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

type Node struct {
	n    int
	next *Node
}

var state *Node
var max int
var findnode []*Node
var curnode *Node

func printstate() {
	cur := state
	for {
		if cur == curnode {
			fmt.Printf("(%d) ", cur.n)
		} else {
			fmt.Printf("%d ", cur.n)
		}
		cur = cur.next
		if cur == state {
			break
		}
	}
	fmt.Printf("\n")
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
		fmt.Printf("pickup %d %d %d\n", rem.n, rem.next.n, rem.next.next.n)
		fmt.Printf("target %d\n", tgt)
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
		fmt.Printf("%d", cur.n)
		cur = cur.next
		if cur == start {
			break
		}
	}
	fmt.Printf("\n")
}

func playgame(v []int, total int, rounds int) {
	state = nil
	findnode = make([]*Node, total+1)
	max = 0

	lastnext := &state

	insert := func(val int) {
		n := &Node{n: val}
		findnode[val] = n
		*lastnext = n
		lastnext = &n.next
		if val > max || max == 0 {
			max = val
		}
	}

	for i := range v {
		insert(v[i])
	}
	for i := max + 1; i < total+1; i++ {
		insert(i)
	}

	*lastnext = state
	curnode = state

	for round := 0; round < rounds; round++ {
		if round%1000 == 0 && round > 0 {
			fmt.Printf("progress %0.04g%%\n", (float64(round)/float64(rounds))*100)
		}
		if debug {
			fmt.Printf("-- move %d --\n", round+1)
			printstate()
		}
		move()
		if debug {
			fmt.Printf("\n")
		}
	}
}

//var input = 389125467 // example
var input = 467528193 // real input

const debug = false

func main() {
	v := vatoi(strings.Split(fmt.Sprintf("%d", input), ""))

	// PART 1

	playgame(v, 9, 100)
	fmt.Printf("final:\n")
	printstate()

	fmt.Printf("PART 1: ")
	printfromlabel1()

	// PART 2
	playgame(v, 1000000, 10000000)

	fmt.Printf("PART 2: %d %d = %d\n", findnode[1].next.n, findnode[1].next.next.n, findnode[1].next.n*findnode[1].next.next.n)
}
