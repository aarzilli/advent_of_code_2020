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
	snd := 0
	regs := map[string]int{}

interpLoop:
	for {
		instr := text[pc]
		switch instr.opcode {
		case "snd":
			snd = instr.args[0].value(regs)
			pc++
		case "set":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[1].value(regs)
			pc++
		case "add":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) + instr.args[1].value(regs)
			pc++
		case "mul":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) * instr.args[1].value(regs)
			pc++
		case "mod":
			instr.argMustBeReg(0)
			regs[instr.args[0].reg] = instr.args[0].value(regs) % instr.args[1].value(regs)
			pc++
		case "rcv":
			if instr.args[0].value(regs) != 0 {
				fmt.Printf("recovered %d\n", snd)
				break interpLoop
			}
			pc++
		case "jgz":
			if instr.args[0].value(regs) > 0 {
				pc += instr.args[1].value(regs)
			} else {
				pc++
			}
		default:
			panic("blah")
		}
	}
}

func main() {
	buf, err := ioutil.ReadFile("../aoc2017bis/18.txt")
	must(err)
	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		text = append(text, parseInstr(line))
	}

	run()
}
