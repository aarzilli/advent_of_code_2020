package main

import "fmt"

var input = []int{20, 0, 1, 11, 6, 3} // real input
//var input = []int{ 0,3,6 } // example

func main() {
	seen := map[int]int{}
	var last int
	var next int

	say := func(num int, round int) {
		//fmt.Printf("SAY %d\n", num)
		oldseen, isseen := seen[num]
		if isseen {
			next = round - oldseen
		} else {
			next = 0
		}
		seen[num] = round
		last = num
	}

	for i := range input {
		say(input[i], i)
	}

	const limit = 30000000

	for round := len(input); round < limit; round++ {
		if round == 2020 {
			fmt.Printf("PART 1: %d\n", last)
		}
		if round%1000000 == 0 {
			fmt.Printf("%0.02g%%\n", float64(round)/float64(limit)*100)
		}
		say(next, round)
	}

	fmt.Printf("PART 2: %d\n", last)
}
