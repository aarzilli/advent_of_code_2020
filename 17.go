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

type Pt struct {
	x, y, z, w int
}

var M = map[Pt]byte{}

func get(x, y, z, w int) byte {
	ch, ok := M[Pt{x, y, z, w}]
	if ok {
		return ch
	}
	return '.'
}

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

func neighbors3d(pt Pt) []Pt {
	r := []Pt{}
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dz == 0 && dy == 0 && dx == 0 {
					continue
				}
				r = append(r, Pt{pt.x + dx, pt.y + dy, pt.z + dz, 0})
			}
		}
	}
	return r
}

func neighbors4d(pt Pt) []Pt {
	r := []Pt{}
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dw == 0 && dz == 0 && dy == 0 && dx == 0 {
						continue
					}
					r = append(r, Pt{pt.x + dx, pt.y + dy, pt.z + dz, pt.w + dw})
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

func step(is4d bool) {
	neighbors := neighbors3d
	if is4d {
		neighbors = neighbors4d
	}

	NewM := make(map[Pt]byte)
	minx, maxx, miny, maxy, minz, maxz, minw, maxw := minmax()

	if !is4d {
		minw = +1
		maxw = -1
	}

	for w := minw - 1; w <= maxw+1; w++ {
		for z := minz - 1; z <= maxz+1; z++ {
			for y := miny - 1; y <= maxy+1; y++ {
				for x := minx - 1; x <= maxx+1; x++ {
					active, _ := count(neighbors(Pt{x, y, z, w}))
					if get(x, y, z, w) == '#' {
						if active == 2 || active == 3 {
							NewM[Pt{x, y, z, w}] = '#'
						}
					} else {
						if active == 3 {
							NewM[Pt{x, y, z, w}] = '#'
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
			M[Pt{x: j, y: i, z: 0}] = line[j]
		}
	}

	OriginalM := M

	for i := 0; i < 6; i++ {
		step(false)
	}
	pf("PART 1: %d\n", total())

	M = OriginalM

	for i := 0; i < 6; i++ {
		step(true)
	}
	pf("PART 2: %d\n", total())
}
