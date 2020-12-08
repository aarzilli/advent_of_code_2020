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

const trace = false

var text []Instr

func run(part1 bool) bool {
	pc := 0
	acc := 0
	seen := map[int]bool{}
	regs := map[string]int{}

	for {
		if seen[pc] {
			if part1 {
				pf("PART1: %d\n", acc)
			}
			return false
		}
		if pc >= len(text) {
			pf("PART 2: %d\n", acc)
			return true
		}
		instr := text[pc]
		seen[pc] = true
		if trace {
			pf("at %d %s %#v\n", pc, text[pc].opcode, text[pc].args)
		}
		switch instr.opcode {
		case "acc":
			acc += instr.args[0].value(regs)
			pc++
		case "nop":
			pc++
		case "jmp":
			pc += instr.args[0].value(regs)
		default:
			panic("blah")
		}
	}
}

func main() {
	buf, err := ioutil.ReadFile("08.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		text = append(text, parseInstr(line))
	}

	run(true)

	for i := range text {
		switch text[i].opcode {
		case "jmp":
			text[i].opcode = "nop"
			if run(false) {
				return
			}
			text[i].opcode = "jmp"
		case "nop":
			text[i].opcode = "jmp"
			if run(false) {
				break
			}
			text[i].opcode = "nop"
		}
	}
}
