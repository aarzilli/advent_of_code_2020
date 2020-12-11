package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

var M [][]byte
var M1 [][]byte

func isocc(i, j int) int {
	if i < 0 || i >= len(M) {
		return 0
	}
	if j < 0 || j >= len(M[i]) {
		return 0
	}
	if M[i][j] == '#' {
		return 1
	}
	return 0
}

func occupied(i, j int) int {
	return isocc(i-1, j-1) + isocc(i-1, j) + isocc(i-1, j+1) + isocc(i, j-1) + isocc(i, j+1) + isocc(i+1, j-1) + isocc(i+1, j) + isocc(i+1, j+1)
}

func stepP1() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'L' && occupied(i, j) == 0 {
				M1[i][j] = '#'
			} else if M[i][j] == '#' && occupied(i, j) >= 4 {
				M1[i][j] = 'L'
			} else {
				M1[i][j] = M[i][j]
			}
		}
	}
	M, M1 = M1, M
}

func nochange() bool {
	for i := range M {
		for j := range M[i] {
			if M[i][j] != M1[i][j] {
				return false
			}
		}
	}
	return true
}

func allocc() int {
	count := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func stepP2() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'L' && occupiedP2(i, j) == 0 {
				M1[i][j] = '#'
			} else if M[i][j] == '#' && occupiedP2(i, j) >= 5 {
				M1[i][j] = 'L'
			} else {
				M1[i][j] = M[i][j]
			}
		}
	}
	M, M1 = M1, M
}

func occupiedP2(i, j int) int {
	return isoccP2(i, j, -1, -1) + isoccP2(i, j, -1, 0) + isoccP2(i, j, -1, +1) + isoccP2(i, j, 0, -1) + isoccP2(i, j, 0, +1) + isoccP2(i, j, +1, -1) + isoccP2(i, j, +1, 0) + isoccP2(i, j, +1, +1)
}

func isoccP2(si, sj, di, dj int) int {
	i := si + di
	j := sj + dj
	for i >= 0 && i < len(M) && j >= 0 && j < len(M[i]) {
		switch M[i][j] {
		case '#':
			return 1
		case 'L':
			return 0
		default:
			i += di
			j += dj
		}
	}
	return 0
}

func printmatrix() {
	for i := range M {
		pf("%s\n", string(M[i]))
	}
	pf("\n")
}

const debug = false

func loop(step func()) {
	for n := 0; n < 100000; n++ {
		//pf("%d %d\n", n, allocc())
		step()
		if debug {
			printmatrix()
		}
		if nochange() {
			break
		}
	}

	if debug {
		printmatrix()
	}
}

func main() {
	buf, err := ioutil.ReadFile("11.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		M = append(M, []byte(line))
	}

	M1 = make([][]byte, len(M))
	for i := range M {
		M1[i] = make([]byte, len(M[i]))
	}

	savedM := make([][]byte, len(M))
	for i := range M {
		savedM[i] = make([]byte, len(M[i]))
		copy(savedM[i], M[i])
	}

	loop(stepP1)

	pf("PART 1: %d\n", allocc())

	M = savedM

	loop(stepP2)

	pf("PART 2: %d\n", allocc())
}
