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
func nolastif(x []string) []string {
	if x[len(x)-1] != "" {
		return x
	}
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

var player1, player2 []int

func round() {
	var winner, loser *[]int
	if player1[0] > player2[0] {
		winner = &player1
		loser = &player2
	} else {
		winner = &player2
		loser = &player1
	}
	
	wincard := (*winner)[0]
	losscard := (*loser)[0]
	(*winner) = (*winner)[1:]
	(*loser) = (*loser)[1:]
	
	(*winner) = append(*winner, wincard)
	(*winner) = append(*winner, losscard)
}

func state() {
	pf("player1=%v\n", player1)
	pf("player2=%v\n\n", player2)
}

func main() {
	buf, err := ioutil.ReadFile("22.txt")
	must(err)
	blocks := strings.Split(string(buf), "\n\n")
	
	player1 = vatoi(nolastif(splitandclean(blocks[0], "\n", -1)[1:]))
	player2 = vatoi(nolastif(splitandclean(blocks[1], "\n", -1)[1:]))
	
	for n := 0; n < 10000; n++ {
		pf("round %d\n", n)
		state()
		round()
		if len(player1) == 0 || len(player2) == 0 {
			break
		}
	}
	
	state()
	
	winner := player1
	if len(winner) == 0 {
		winner = player2
	}
	
	tot := 0
	for i := range winner {
		tot += winner[i] * (len(winner)-i)
		pf("%d %d\n", winner[i], (len(winner)-i))
	}
	pf("PART 1: %d\n", tot)
}
