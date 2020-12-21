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

type Food struct {
	Ingrs map[string]bool
	Alls map[string]bool
}

var foods []*Food
var allergenes = map[string]bool{}

func main() {
	buf, err := ioutil.ReadFile("21.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := splitandclean(line, "(contains", -1)
		ings := splitandclean(v[0], " ", -1)
		alls := splitandclean(nolast(v[1]), ",", -1)
		pf("%q %q\n", ings, alls)
		
		var food Food
		food.Ingrs = make(map[string]bool)
		food.Alls = make(map[string]bool)
		
		for _, s := range ings {
			food.Ingrs[s] = true
		}
		for _, s := range alls {
			food.Alls[s] = true
			allergenes[s] = true
		}
		
		foods = append(foods, &food)
	}
	
	possiblyunsafe := map[string]bool{}
	
	for all := range allergenes {
		pf("allergene %s\n", all)
		ingrs := make(map[string]bool)
		first := true
		for _, food := range foods {
			if !food.Alls[all] {
				continue
			}
			if first {
				first = false
				for ingr := range food.Ingrs {
					ingrs[ingr] = true
				}
			} else {
				for ingr := range ingrs {
					if !food.Ingrs[ingr] {
						delete(ingrs, ingr)
					}
				}
			}
			pf("\t%#v\n", food.Ingrs)
			pf("\tcur %#v\n", ingrs)
		}
		
		pf("\tfinal %#v\n", ingrs)
		for ingr := range ingrs {
			possiblyunsafe[ingr] = true
		}
	}
	
	pf("\npossibly unsafe: %#v\n", possiblyunsafe)
	
	n := 0
	for _, food := range foods {
		for ingr := range food.Ingrs {
			if !possiblyunsafe[ingr] {
				pf("%s\n", ingr)
				n++
			}
		}
	}
	
	pf("PART 1: %d\n", n)
}
