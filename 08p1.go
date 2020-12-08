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

type Instr struct {
	opcode string
	args   []Arg
}

func parseInstr(line string) Instr {
	fields := splitandclean(line, " ", -1)
	opcode := fields[0]
	args := make([]Arg, len(fields)-1)
	for i, field := range fields[1:] {
		if field[len(field)-1] == ',' {
			field = field[:len(field)-1]
		}
		n, err := strconv.Atoi(field)
		if err == nil {
			args[i] = Arg{val: n}
		} else {
			args[i] = Arg{reg: field}
		}
	}
	return Instr{opcode, args}
}

func (instr Instr) argMustBeReg(argnum int) {
	if instr.args[argnum].reg == "" {
		panic("arg is not register")
	}
}

type Arg struct {
	reg string
	val int
}

func (a Arg) value(regs map[string]int) int {
	if a.reg == "" {
		return a.val
	}
	return regs[a.reg]
}

var text []Instr

func run() {
	pc := 0
	acc := 0
	seen := map[int]bool{}

	for {
		if seen[pc] {
			pf("PART1: %d\n", acc)
			break
		}
		instr := text[pc]
		seen[pc] = true
		pf("at %d %#v\n", pc, text[pc])
		switch instr.opcode {
		case "acc":
			acc += instr.args[0].val
			pc++
		case "nop":
			pc++
		case "jmp":
			pc += instr.args[0].val
		default:
			panic("blah")
		}
	}
}

func main() {
	buf, err := ioutil.ReadFile("08.example2.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		text = append(text, parseInstr(line))
	}
	
	run()
}
