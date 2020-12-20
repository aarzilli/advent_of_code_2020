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

func atoi(in string) int {
	n, err := strconv.Atoi(in)
	must(err)
	return n
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

func copytile(tile *Tile) *Tile {
	var r Tile
	r.id = tile.id
	r.M = make([][]byte, len(tile.M))
	for i := range tile.M {
		r.M[i] = append(r.M[i], tile.M[i]...)
	}
	return &r
}

func views(tile *Tile) []*Tile {
	r := []*Tile{tile}

	flips := func(tile *Tile) {
		t2 := copytile(tile)
		change(t2, true, false, false)
		r = append(r, t2)

		t2 = copytile(tile)
		change(t2, false, true, false)
		r = append(r, t2)

		t2 = copytile(tile)
		change(t2, true, true, false)
		r = append(r, t2)
	}

	flips(tile)

	trot := copytile(tile)
	change(trot, false, false, true)
	r = append(r, trot)
	flips(trot)

	return r
}

var arrang [][]*Tile

const debug = false

func findnext(i, j, myb, otherb int) *Tile {
	t := arrang[i][j]
	b := borders(t)

	if debug {
		pf("%q\n", b)
	}

	myborder := b[myb]

	if len(assoc[myborder]) > 2 {
		panic(fmt.Errorf("can't deal with this %#v", assoc[myborder]))
	}

	for _, c := range assoc[myborder] {
		if c.id == t.id {
			continue
		}

		t2 := gettile(c.id)

		for _, t3 := range views(t2) {
			if borders(t3)[otherb] == myborder {
				return t3
			}
		}

		panic("blah")
	}

	return nil
}

func printile(t *Tile) {
	for i := range t.M {
		pf("%s\n", t.M[i])
	}
}

func getcoord(tile *Tile, i, j int) byte {
	if i < 0 || i >= len(tile.M) {
		return '.'
	}
	if j < 0 || j >= len(tile.M[i]) {
		return '.'
	}
	return tile.M[i][j]
}

var monster = [][]byte{
	[]byte(".#...#.###...#.##.O#.."),
	[]byte("O.##.OO#.#.OO.##.OOO##"),
	[]byte("#O.#O#.O##O..O.#O##.##"),
}

const debugmonster = false

func findmonster(bigm *Tile, si, sj int) bool {
	for i := range monster {
		for j := range monster[i] {
			if monster[i][j] != 'O' {
				continue
			}
			if getcoord(bigm, si+i, sj+j) != '#' {
				return false
			}
		}
	}

	if debugmonster {
		pf("found at %d,%d\n", si, sj)
	}

	for i := range monster {
		for j := range monster[i] {
			if monster[i][j] == 'O' {
				bigm.M[si+i][sj+j] = 'O'
			}
		}
	}
	return true
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

	m := map[int]int{}

	for _, v := range assoc {
		if len(v) > 1 {
			for _, c := range v {
				m[c.id]++
			}
		}
	}

	minid := 0

	tot := 1
	for id := range m {
		if m[id] == 4 {
			if minid == 0 || id < minid {
				minid = id
			}
			tot *= id
		}
	}

	pf("PART 1: %d\n", tot)

	if file == "20.example.txt" {
		minid = 1951
	}

	arrang = make([][]*Tile, 200)
	for i := range arrang {
		arrang[i] = make([]*Tile, 200)
	}

	// pick the smallest corner, then rotate and flip it until it can fit as the top left corner
	for _, t := range views(gettile(minid)) {
		b := borders(t)
		if len(assoc[b[1]]) == 2 && len(assoc[b[3]]) == 2 {
			arrang[0][0] = t
			break
		}
	}

	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			//pf("%d,%d\n", i, j)
			if arrang[i][j] == nil {
				break
			}
			arrang[i+1][j] = findnext(i, j, 1, 0)
			arrang[i][j+1] = findnext(i, j, 3, 2)
		}
	}

	maxi, maxj := 0, 0

	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			if arrang[i][j] == nil {
				break
			}
			maxi = i
			maxj = j

			tile := arrang[i][j]
			tile.M = tile.M[1:]
			tile.M = tile.M[:len(tile.M)-1]
			for i := range tile.M {
				tile.M[i] = tile.M[i][1:]
				tile.M[i] = tile.M[i][:len(tile.M[i])-1]
			}
		}
	}

	bigm := make([][]byte, 1)

	for i := 0; i <= maxi; i++ {
		t := arrang[i][0]
		for row := 0; row < len(t.M); row++ {
			for j := 0; j <= maxj; j++ {
				t := arrang[i][j]
				bigm[len(bigm)-1] = append(bigm[len(bigm)-1], t.M[row]...)
			}
			bigm = append(bigm, []byte{})
		}
	}

	if len(bigm[len(bigm)-1]) == 0 {
		bigm = bigm[:len(bigm)-1]
	}

	for _, t := range views(&Tile{M: bigm}) {
		ok := false
		for i := range t.M {
			for j := range t.M[i] {
				if findmonster(t, i, j) {
					ok = true
				}
			}
		}
		if ok {
			tot2 := 0
			for i := range t.M {
				for j := range t.M[i] {
					if t.M[i][j] == '#' {
						tot2++
					}
				}
			}

			pf("PART 2: %d\n", tot2)
		}
	}
}
