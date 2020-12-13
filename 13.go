package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"math/big"
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
		if err != nil {
			r[i] = -1
		}
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

var one = big.NewInt(1)
 
func crt(a, n []*big.Int) (*big.Int, error) {
	// stolen from rosetta code :)
    p := new(big.Int).Set(n[0])
    for _, n1 := range n[1:] {
        p.Mul(p, n1)
    }
    var x, q, s, z big.Int
    for i, n1 := range n {
        q.Div(p, n1)
        z.GCD(nil, &s, n1, &q)
        if z.Cmp(one) != 0 {
            return nil, fmt.Errorf("%d not coprime", n1)
        }
        x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
    }
    return x.Mod(&x, p), nil
}

func main() {
	buf, err := ioutil.ReadFile("13.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	start := atoi(lines[0])
	buses := vatoi(splitandclean(lines[1], ",", -1))
	pf("%d %#v\n", start, buses)
	
	min := 0
	for i := range buses {
		if buses[i] > min {
			min = buses[i]
		}
	}
	min++
	minid := -1
	
	for _, id := range buses {
		if id < 0 {
			continue
		}
		
		pos := start % id
		if pos == 0 {
			minid = id
			min = 0
		} else {
			diff := id - pos
			if diff < min {
				minid = id
				min = diff
			}
		}
	}
	
	pf("PART 1: %d %d %d\n", min, minid, min*minid)
	
	a := make([]*big.Int, 0)
	n := make([]*big.Int, 0)
	
	for i, id := range buses {
		if id < 0 {
			continue
		}
		a = append(a, big.NewInt(int64(id-i)))
		n = append(n, big.NewInt(int64(id)))
	}
	
	pf("%v %v\n", a, n)
	
	sol, err := crt(a, n)
	pf("PART 2: %v %v\n", sol, err)
}
