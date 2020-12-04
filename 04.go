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

type Passport struct {
	v []string
}

var passports []*Passport


func (p *Passport) valid() bool {
	fields := map[string]bool{}
	for i := range p.v {
		v := splitandclean(p.v[i], ":", -1)
		fields[v[0]] = true
	}
	
	for _, m := range []string{ "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" } {
		if !fields[m] {
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

func (p *Passport) valid2() bool {
	if !p.valid() {
		return false
	}
	fields := map[string]string{}
	for i := range p.v {
		v := splitandclean(p.v[i], ":", -1)
		fields[v[0]] = v[1]
	}
	
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
	fmt.Printf("hello\n")
	buf, err := ioutil.ReadFile("04.txt")
	must(err)
	
	cur := &Passport{}
	
	add := func(line string) {
		v := splitandclean(line, " ", -1)
		cur.v = append(cur.v, v...)
	}
	
	flush := func() {
		if len(cur.v) != 0 {
			passports = append(passports, cur)
			cur = &Passport{}
		}
	}
	
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			flush()
			continue
		}
		
		add(line)
	}
	
	flush()
	
	n, n2 := 0, 0
	for i := range passports {
		if passports[i].valid() {
			n++
		}
		if passports[i].valid2() {
			n2++
		}
		if !passports[i].valid() && passports[i].valid2() {
			fmt.Printf("error!\n")
		}
	}
	
	fmt.Printf("PART 1: %d\n", n)
	fmt.Printf("PART 2: %d\n", n2)
}
