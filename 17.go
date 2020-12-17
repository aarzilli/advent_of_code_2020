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

type Pt struct {
	x,y,z,w int
}

var M = map[Pt]byte{}

func get(x, y, z, w int) byte {
	ch, ok := M[Pt{x,y,z, w}]
	if ok {
		return ch
	}
	return '.'
}

const MAX = 1000000000

func minmax() (minx, maxx, miny, maxy, minz, maxz int, minw, maxw int) {
	first := true
	for pt, _ := range M {
		if first {
			minx = pt.x
			maxx = pt.x
			miny = pt.y
			maxy = pt.y
			minz = pt.z
			maxx = pt.z
			minw = pt.w
			maxw = pt.w
			first = false
		}
		if pt.x < minx {
			minx = pt.x
		}
		if pt.x > maxx {
			maxx = pt.x
		}
		if pt.y < miny {
			miny = pt.y
		}
		if pt.y > maxy {
			maxy = pt.y
		}
		if pt.z < minz {
			minz = pt.z
		}
		if pt.z > maxz {
			maxz = pt.z
		}
		if pt.w < minw {
			minw = pt.w
		}
		if pt.w > maxw {
			maxw = pt.w
		}
	}
	return
}

func printcube() {
	minx, maxx, miny, maxy, minz, maxz, minw, maxw := minmax()
	
	_, _ = minw, maxw // TBD
	
	for z := minz; z <= maxz; z++ {
		pf("z=%d\n", z)
		for y := miny; y <= maxy; y++ {
			for x := minx; x <= maxx; x++ {
				pf("%c", get(x, y, z, 0))
			}
			pf("\n")
		}
		pf("\n")
	}
}

func neighbors(pt Pt) []Pt {
	r := []Pt{}
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dw == 0 && dz == 0 && dy == 0 && dx == 0 {
						continue
					}
					r = append(r, Pt{pt.x+dx, pt.y+dy, pt.z+dz, pt.w+dw})
				}
			}
		}
	}
	return r
}

func count(pts []Pt) (active, inactive int) {
	for _, pt := range pts {
		if get(pt.x, pt.y, pt.z, pt.w) == '.' {
			inactive++
		} else {
			active++
		}
	}
	return
}

func step() {
	NewM := make(map[Pt]byte)
	minx, maxx, miny, maxy, minz, maxz, minw, maxw := minmax()
	
	for w := minw-1; w <= maxw+1; w++ {
	for z := minz-1; z <= maxz+1; z++ {
		for y := miny-1; y <= maxy+1; y++ {
			for x := minx-1; x <= maxx+1; x++ {
				active, _ := count(neighbors(Pt{x,y,z,w}))
				if get(x, y, z, w) == '#' {
					if active == 2 || active == 3 {
						NewM[Pt{x,y,z,w}] = '#'
					}
				} else {
					if active == 3 {
						NewM[Pt{x,y,z,w}] = '#'
					}
				}
			}
		}
	}
	}
	M = NewM
}

func total() int {
	r := 0
	for _, ch := range M {
		if ch == '#' {
			r++
		}
	}
	return r
}

func main() {
	buf, err := ioutil.ReadFile("17.txt")
	must(err)
	lines := strings.Split(string(buf), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for j := range line {
			M[Pt{ x: j, y: i, z: 0 }] = line[j]
		}
	}
	
	for i := 0; i < 6; i++ {
		step()
	}
	
	pf("PART 1: %d\n", total())
}
