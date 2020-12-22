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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

const (
	debug1 = false
	debug2 = false
)

type Game struct {
	id                 int
	cnt                int
	player1, player2   []int
	p1states, p2states map[string]bool
}

var gameid = 0

func newgame(player1, player2 []int) *Game {
	g := &Game{}
	g.id = gameid
	gameid++
	g.player1 = append(g.player1, player1...)
	g.player2 = append(g.player2, player2...)
	g.p1states = make(map[string]bool)
	g.p2states = make(map[string]bool)
	return g
}

func (g *Game) state() {
	pf("player1=%v\n", g.player1)
	pf("player2=%v\n", g.player2)
}

func resolve(winner, loser *[]int) {
	// this leaks memory but it doesn't matter
	wincard := (*winner)[0]
	losscard := (*loser)[0]
	(*winner) = (*winner)[1:]
	(*loser) = (*loser)[1:]
	(*winner) = append(*winner, wincard)
	(*winner) = append(*winner, losscard)
}

func (g *Game) playgame1() {
	for n := 0; true; n++ {
		if debug1 {
			pf("round %d\n", n)
			g.state()
		}
		if g.player1[0] > g.player2[0] {
			resolve(&g.player1, &g.player2)
		} else {
			resolve(&g.player2, &g.player1)
		}
		if len(g.player1) == 0 || len(g.player2) == 0 {
			break
		}
	}
}

func checkrepeat(deck []int, seen map[string]bool) bool {
	state := fmt.Sprintf("%v", deck)
	if seen[state] {
		return true
	}
	seen[state] = true
	return false
}

func (g *Game) round2() bool {
	if debug2 {
		pf("-- Round %d (Game %d) --\n", g.cnt+1, g.id+1)
		g.state()
	}
	if checkrepeat(g.player1, g.p1states) || checkrepeat(g.player2, g.p2states) {
		if debug2 {
			pf("end by check\n")
		}
		return false
	}

	if g.player1[0] < len(g.player1) && g.player2[0] < len(g.player2) {
		if debug2 {
			pf("recursive!\n")
		}
		g2 := newgame(g.player1[1:][:g.player1[0]], g.player2[1:][:g.player2[0]])
		winid := g2.playgame2()
		if winid == 1 {
			resolve(&g.player1, &g.player2)
		} else {
			resolve(&g.player2, &g.player1)
		}
	} else {
		if g.player1[0] > g.player2[0] {
			resolve(&g.player1, &g.player2)
			if debug2 {
				pf("player1 wins (round %d game %d)!\n", g.cnt, g.id+1)
			}
		} else {
			resolve(&g.player2, &g.player1)
			if debug2 {
				pf("player2 wins (round %d game %d)!\n", g.cnt, g.id+1)
			}
		}
	}

	if debug2 {
		pf("\n")
	}
	g.cnt++

	return true
}

var cache = map[string]int{}

func (g *Game) playgame2() int {
	key := fmt.Sprintf("player1=%v player2=%v", g.player1, g.player2)
	if r, ok := cache[key]; ok {
		return r
	}
	for {
		ok := g.round2()
		if !ok {
			cache[key] = 1
			return 1
		}
		if len(g.player1) == 0 || len(g.player2) == 0 {
			break
		}
	}

	var r int
	if len(g.player1) == 0 {
		r = 2
	} else {
		r = 1
	}

	cache[key] = r
	return r
}

func calcscore(winner []int) int {
	tot := 0
	for i := range winner {
		tot += winner[i] * (len(winner) - i)
	}
	return tot
}

func main() {
	buf, err := ioutil.ReadFile("22.txt")
	must(err)
	blocks := strings.Split(string(buf), "\n\n")

	player1 := vatoi(nolastif(splitandclean(blocks[0], "\n", -1)[1:]))
	player2 := vatoi(nolastif(splitandclean(blocks[1], "\n", -1)[1:]))

	g := newgame(player1, player2)
	g.playgame1()
	pf("PART 1: %d %d\n", calcscore(g.player1), calcscore(g.player2))

	g = newgame(player1, player2)
	g.playgame2()
	if debug2 {
		g.state()
	}
	pf("PART 2: %d %d\n", calcscore(g.player1), calcscore(g.player2))
}
