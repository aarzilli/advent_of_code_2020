package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
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

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

type Food struct {
	Ingrs map[string]bool
	Alls  map[string]bool
}

var foods []*Food
var allergenes = map[string]bool{}
var mapping = map[string]string{} // allergene -> ingredient

func eliminateallergene() bool {
	for all := range allergenes {
		if mapping[all] != "" {
			continue
		}
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
		}

		if len(ingrs) == 1 {
			for ingr := range ingrs {
				pf("found allergene %s as %s\n", all, ingr)
				mapping[all] = ingr
				break
			}

			ingr := mapping[all]
			for _, food := range foods {
				if food.Ingrs[ingr] {
					delete(food.Ingrs, ingr)
				}
			}
			return true
		}
	}
	return false
}

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
		}

		for ingr := range ingrs {
			possiblyunsafe[ingr] = true
		}
	}

	n := 0
	for _, food := range foods {
		for ingr := range food.Ingrs {
			if !possiblyunsafe[ingr] {
				delete(food.Ingrs, ingr)
				n++
			}
		}
	}

	pf("PART 1: %d\n", n)

	for {
		if !eliminateallergene() {
			break
		}
	}

	alls := []string{}
	for all := range allergenes {
		alls = append(alls, all)
	}
	sort.Strings(alls)
	ingrs := []string{}
	for _, all := range alls {
		ingrs = append(ingrs, mapping[all])
	}

	pf("PART 2: %s\n", strings.Join(ingrs, ","))
}
