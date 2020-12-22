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

type Game struct {
	id int
	cnt int
	player1, player2 []int
	p1states, p2states map[string]bool
}

func (g *Game) state() {
	pf("player1=%v\n", g.player1)
	pf("player2=%v\n", g.player2)
}

func (g *Game) round() bool {
	//doprint := g.id < 10
	doprint := false
	if doprint {	
		pf("-- Round %d (Game %d) --\n", g.cnt+1, g.id+1)
		g.state()
	}
	p1state := fmt.Sprintf("%v", g.player1)
	p2state := fmt.Sprintf("%v", g.player2)
	if g.p1states[p1state] {
		if doprint {
			pf("end by check\n")
		}
		return false
	}
	if g.p2states[p2state] {
		if doprint {
			pf("end by check\n")
		}
		return false
	}
	
	g.p1states[p1state] = true
	g.p2states[p2state] = true
	
	var winner, loser *[]int
	
	if g.player1[0] < len(g.player1) && g.player2[0] < len(g.player2) {
		if doprint {
			pf("recursive!\n")
		}
		g2 := newgame(g.player1[1:][:g.player1[0]], g.player2[1:][:g.player2[0]])
		winid := g2.playgame()
		if winid == 1 {
			winner = &g.player1
			loser = &g.player2
		} else {
			pf("PLAYER 2 WON?!\n")
			winner = &g.player2
			loser = &g.player1
		}
	} else {
		if g.player1[0] > g.player2[0] {
			winner = &g.player1
			loser = &g.player2
			if doprint {
				pf("player1 wins (round %d game %d)!\n", g.cnt, g.id+1)
			}
		} else {
			winner = &g.player2
			loser = &g.player1
			if doprint {
				pf("player2 wins (round %d game %d)!\n", g.cnt, g.id+1)
			}
		}
	}
	
	wincard := (*winner)[0]
	losscard := (*loser)[0]
	(*winner) = (*winner)[1:]
	(*loser) = (*loser)[1:]
	
	(*winner) = append(*winner, wincard)
	(*winner) = append(*winner, losscard)
	
	if doprint {
		pf("\n")
	}
	g.cnt++
	
	return true
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

var cache = map[string]int{}

func (g *Game) playgame() int {
	key := fmt.Sprintf("player1=%v player2=%v", g.player1, g.player2)
	if r, ok := cache[key]; ok {
		pf("skipping game %d, player %d will win\n", g.id, r)
		return r
	}
	for {
		ok := g.round()
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
		tot += winner[i] * (len(winner)-i)
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
	
	winner := g.playgame()
	pf("winner is %d\n", winner)
	
	pf("post game:\n")
	g.state()
	pf("PART 2: %d %d\n", calcscore(g.player1), calcscore(g.player2))
}
