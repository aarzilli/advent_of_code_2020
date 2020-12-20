package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

var tiles = []*Tile{}

type Tile struct {
	id int
	M  [][]byte
	n  int
}

func borders(tile *Tile) []string {
	r := []string{}
	r = append(r, string(tile.M[0]))
	r = append(r, string(tile.M[len(tile.M)-1]))

	left := []byte{}
	right := []byte{}
	for i := range tile.M {
		left = append(left, tile.M[i][0])
		right = append(right, tile.M[i][len(tile.M[i])-1])
	}

	r = append(r, string(left))
	r = append(r, string(right))
	return r
}

func reverse(s string) string {
	r := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return string(r)
}

var assoc = map[string][]conf{}

type conf struct {
	borderi int
	flip    bool
	id      int
}

func contains(v []conf, id int) (bool, int, bool) {
	for i := range v {
		if v[i].id == id {
			return true, v[i].borderi, v[i].flip
		}
	}
	return false, -1, false
}

type conf2 struct {
	id    int
	fliph bool
	flipv bool
}

func gettile(id int) *Tile {
	for _, tile := range tiles {
		if tile.id == id {
			return tile
		}
	}
	panic("blah")
}

func change(tile *Tile, fliph, flipv, rotate bool) {
	if flipv {
		for i := range tile.M {
			tile.M[i] = []byte(reverse(string(tile.M[i])))
		}
	}

	if fliph {
		newm := make([][]byte, 0, len(tile.M))
		for i := len(tile.M) - 1; i >= 0; i-- {
			newm = append(newm, tile.M[i])
		}
		tile.M = newm
	}
	if rotate {
		newm := make([][]byte, len(tile.M))
		for i := range newm {
			newm[i] = make([]byte, len(tile.M[0]))
		}

		for i := range tile.M {
			for j := range tile.M[i] {
				newm[i][j] = tile.M[len(tile.M)-j-1][i]
			}
		}
		tile.M = newm
	}
}

var arrang [][]int

const debug = false

func findnext(i, j int) {
	id := arrang[i][j]
	t := gettile(id)
	b := borders(t)

	if debug {
		pf("%q\n", b)
	}

	bottom := b[1]
	right := b[3]

	if len(assoc[bottom]) > 2 {
		panic(fmt.Errorf("can't deal with this %#v", assoc[bottom]))
	}

	if len(assoc[right]) > 2 {
		panic(fmt.Errorf("can't deal with this %#v", assoc[right]))
	}

	for _, c := range assoc[bottom] {
		if c.id == id {
			continue
		}

		if debug {
			pf("%d,%d is %d\n", i+1, j, c.id)
		}
		arrang[i+1][j] = c.id

		t2 := gettile(c.id)

		if borders(t2)[0] != bottom {
			switch c.borderi {
			case 0:
				if c.flip {
					change(t2, false, true, false)
				} else {
					// ok as is
				}
			case 1:
				if c.flip {
					change(t2, true, true, false)
				} else {
					change(t2, true, false, false)
				}
			case 2:
				change(t2, false, false, true)
				change(t2, false, !c.flip, false)
			case 3:
				change(t2, false, false, true)
				change(t2, true, !c.flip, false)
			}
		}

		if borders(t2)[0] != bottom {
			printile(t2)
			pf("%#v\n", c)
			pf("%q %q\n", borders(t2)[0], bottom)
			panic("mismatch")
		}

		if debug {
			printile(t2)
		}

		break
	}

	for _, c := range assoc[right] {
		if c.id == id {
			continue
		}

		if debug {
			pf("%d,%d is %d\n", i, j+1, c.id)
		}
		arrang[i][j+1] = c.id

		t2 := gettile(c.id)

		if borders(t2)[2] != right {
			switch c.borderi {
			case 0:
				change(t2, false, false, true)
				change(t2, c.flip, true, false)
			case 1:
				change(t2, false, false, true)
				change(t2, c.flip, false, false)
			case 2:
				if c.flip {
					change(t2, true, false, false)
				} else {
					// of as is
				}
			case 3:
				if c.flip {
					change(t2, true, true, false)
				} else {
					change(t2, false, true, false)
				}
			}
		}

		if borders(t2)[2] != right {
			pf("%#v\n", c)
			pf("%q %q\n", borders(t2)[2], right)
			panic("mismatch")
		}

		if debug {
			printile(t2)
		}

		break
	}
}

func printile(t *Tile) {
	for i := range t.M {
		pf("%s\n", t.M[i])
	}
}

var bigm [][]byte

var monster = [][]byte{
	[]byte(".#...#.###...#.##.O#.."),
	[]byte("O.##.OO#.#.OO.##.OOO##"),
	[]byte("#O.#O#.O##O..O.#O##.##"),
}

func getbigm(i, j int) byte {
	if i < 0 || i >= len(bigm) {
		return '.'
	}
	if j < 0 || j >= len(bigm[i]) {
		return '.'
	}
	return bigm[i][j]
}

func findmonster(si, sj int) {
	for i := range monster {
		for j := range monster[i] {
			if monster[i][j] != 'O' {
				continue
			}
			if getbigm(si+i, sj+j) != '#' {
				return
			}
		}
	}

	pf("found at %d,%d\n", si, sj)

	for i := range monster {
		for j := range monster[i] {
			if monster[i][j] == 'O' {
				bigm[si+i][sj+j] = 'O'
			}
		}
	}
}

func main() {
	//file := "20.example.txt"
	file := "20.txt"
	buf, err := ioutil.ReadFile(file)
	must(err)
	blocks := strings.Split(string(buf), "\n\n")
	for _, block := range blocks {
		cur := &Tile{}

		for _, line := range splitandclean(block, "\n", -1) {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if strings.HasPrefix(line, "Tile ") {
				cur.id = atoi(nolast(splitandclean(line, " ", -1)[1]))
			} else {
				cur.M = append(cur.M, []byte(line))
			}
		}

		tiles = append(tiles, cur)
	}

	for _, tile := range tiles {
		for i, border := range borders(tile) {
			assoc[border] = append(assoc[border], conf{i, false, tile.id})
			revb := reverse(border)
			assoc[revb] = append(assoc[revb], conf{i, true, tile.id})
		}
	}

	//pf("%#v\n", assoc)

	m := map[int]int{}

	for _, v := range assoc {
		if len(v) > 1 {
			for _, c := range v {
				m[c.id]++
			}
			//pf("%#v\n", v)
		}
	}

	minid := 0

	corners := map[int]bool{}

	tot := 1
	for id := range m {
		if m[id] == 4 {
			if minid == 0 || id < minid {
				minid = id
			}
			corners[id] = true
			pf("%d\n", id)
			tot *= id
		}
	}

	pf("PART 1: %d\n", tot)

	if file == "20.example.txt" {
		minid = 1951
	}

	arrang = make([][]int, 200)
	for i := range arrang {
		arrang[i] = make([]int, 200)
	}

	arrang[0][0] = minid
	fliph := false
	flipv := false

	for _, v := range assoc {
		if len(v) == 1 {
			continue
		}

		ok, meborder, _ := contains(v, minid)
		if ok {
			if meborder == 0 {
				fliph = true
			}
			if meborder == 2 {
				flipv = true
			}
			//pf("%d\n", meborder)
		}
	}

	_, _ = fliph, flipv

	pf("%v %v\n", fliph, flipv)

	change(gettile(minid), fliph, flipv, false)

	pf("%d %v\n", minid, arrang[0][0])

	/*t := gettile(minid)
	printile(t)*/

	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			//pf("%d,%d\n", i, j)
			if arrang[i][j] == 0 {
				break
			}
			findnext(i, j)
		}
	}

	maxi, maxj := 0, 0

	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			if arrang[i][j] == 0 {
				break
			}
			maxi = i
			maxj = j
		}
	}

	for _, tile := range tiles {
		tile.M = tile.M[1:]
		tile.M = tile.M[:len(tile.M)-1]
		for i := range tile.M {
			tile.M[i] = tile.M[i][1:]
			tile.M[i] = tile.M[i][:len(tile.M[i])-1]
		}

	}

	bigm = make([][]byte, 1)

	for i := 0; i <= maxi; i++ {
		t := gettile(arrang[i][0])
		for row := 0; row < len(t.M); row++ {
			for j := 0; j <= maxj; j++ {
				t := gettile(arrang[i][j])
				bigm[len(bigm)-1] = append(bigm[len(bigm)-1], t.M[row]...)
			}
			bigm = append(bigm, []byte{})
		}
	}

	if len(bigm[len(bigm)-1]) == 0 {
		bigm = bigm[:len(bigm)-1]
	}

	bigtile := &Tile{M: bigm}
	change(bigtile, true, false, false)
	//change(bigtile, false, false, true)
	//change(bigtile, false, true, false)
	//printile(bigtile)
	//pf("\n")
	bigm = bigtile.M

	for i := range bigm {
		for j := range bigm[i] {
			findmonster(i, j)
		}
	}

	tot2 := 0
	for i := range bigm {
		for j := range bigm[i] {
			if bigm[i][j] == '#' {
				tot2++
			}
		}
	}

	pf("PART 2: %d\n", tot2)

}
