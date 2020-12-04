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

// convert string to integer
func atoi(in string) int {
	n, err := strconv.Atoi(in)
	must(err)
	return n
}

func main() {
	buf, err := ioutil.ReadFile("01.txt")
	must(err)
	v := []int{}
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v = append(v, atoi(line))
	}
part1:
	for i := range v {
		for j := i + 1; j < len(v); j++ {
			if v[i]+v[j] == 2020 {
				fmt.Printf("PART 1: %d\n", v[i]*v[j])
				break part1
			}
		}
	}

part2:
	for i := range v {
		for j := i + 1; j < len(v); j++ {
			for k := j + 1; k < len(v); k++ {
				if v[i]+v[j]+v[k] == 2020 {
					fmt.Printf("PART 2: %d\n", v[i]*v[j]*v[k])
					break part2
				}
			}
		}
	}
}
