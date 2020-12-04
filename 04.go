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

func valid1(p map[string]string) bool {
	for _, m := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p[m]; !ok {
			return false
		}
	}

	return true
}

func validDate(dt string, start, end int) bool {
	if len(dt) != 4 {
		return false
	}

	n := atoi(dt)
	if n < start || n > end {
		return false
	}
	return true
}

func validHeight(hgt string) bool {
	if !(strings.HasSuffix(hgt, "cm") || strings.HasSuffix(hgt, "in")) {
		return false
	}
	n := atoi(hgt[:len(hgt)-2])

	if strings.HasSuffix(hgt, "cm") {
		if n < 150 || n > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		if n < 59 || n > 76 {
			return false
		}
	}
	return true
}

func validHair(hair string) bool {
	if hair == "" {
		return false
	}
	if hair[0] != '#' {
		return false
	}

	for _, ch := range hair[1:] {
		if ((ch < '0') || (ch > '9')) && ((ch < 'a') || (ch > 'z')) {
			return false
		}
	}
	return true
}

var validEye = map[string]bool{
	"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
}

func validPid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	for _, ch := range pid {
		if (ch < '0') || (ch > '9') {
			return false
		}
	}
	return true
}

func valid2(fields map[string]string) bool {
	if !validDate(fields["byr"], 1920, 2002) {
		return false
	}

	if !validDate(fields["iyr"], 2010, 2020) {
		return false
	}

	if !validDate(fields["eyr"], 2020, 2030) {
		return false
	}

	if !validHeight(fields["hgt"]) {
		return false
	}

	if !validHair(fields["hcl"]) {
		return false
	}

	if !validEye[fields["ecl"]] {
		return false
	}

	if !validPid(fields["pid"]) {
		return false
	}

	return true
}

func main() {
	buf, err := ioutil.ReadFile("04.txt")
	must(err)

	n1, n2 := 0, 0
	for _, passtr := range splitandclean(string(buf), "\n\n", -1) {
		cur := map[string]string{}
		for _, line := range splitandclean(passtr, "\n", -1) {
			for _, field := range splitandclean(line, " ", -1) {
				kv := splitandclean(field, ":", -1)
				cur[kv[0]] = kv[1]
			}
		}

		if valid1(cur) {
			n1++
			if valid2(cur) {
				n2++
			}
		}
	}

	pf("PART 1: %d\n", n1)
	pf("PART 2: %d\n", n2)
}
