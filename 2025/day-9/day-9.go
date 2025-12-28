package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Grid struct {
	items    [][]byte
	min      *Point
	max      *Point
	redTiles []*Point
}

type Rect struct {
	x      int
	y      int
	width  int
	height int
	size   int
	lower  Point
	upper  Point
}

type Point struct {
	x, y int
}

const SPOT_EMPTY = '.'
const SPOT_RED = '#'
const SPOT_TAKEN = '0'

// read each line in a file as an battery
func ReadInput(path string) Grid {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	grid := Grid{items: make([][]byte, 0), redTiles: make([]*Point, 0)}
	grid.min = &Point{x: math.MaxInt, y: math.MaxInt}
	grid.max = &Point{x: math.MinInt, y: math.MinInt}
	r := bufio.NewReader(f)
	for {
		s, e := r.ReadString('\n')
		if len(s) > 0 {
			endIndex := len(s) - 1
			// last line won't have \n if it's not there...
			// so don't skip last character
			if e == io.EOF {
				endIndex++
			}

			coords := strings.Split(s[0:endIndex], ",")
			x, xerr := strconv.Atoi(coords[0])
			y, yerr := strconv.Atoi(coords[1])
			if xerr != nil || yerr != nil {
				log.Panicln("Could not parse number:", s[0:endIndex], xerr, yerr)
			}
			p := &Point{x: x, y: y}
			grid.redTiles = append(grid.redTiles, p)
			if grid.max.x < p.x {
				grid.max.x = p.x
			}
			if grid.max.y < p.y {
				grid.max.y = p.y
			}
			if grid.min.x > p.x {
				grid.min.x = p.x
			}
			if grid.min.y > p.y {
				grid.min.y = p.y
			}

		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return grid
}

func PrintGrid(items [][]byte) {
	for y := range items {
		b := strings.Builder{}
		for x := range items[y] {
			b.WriteRune(rune(items[y][x]))
		}
		log.Println(b.String())
	}
}

func CopyGrid(grid *Grid) *Grid {
	ng := Grid{items: make([][]byte, len(grid.items))}
	for y := range grid.items {
		ng.items[y] = make([]byte, len(ng.items[y]))
		copy(ng.items[y], grid.items[y])
	}

	return &ng
}

func MakeRect(lower, upper *Point) *Rect {
	r := Rect{}

	if lower.x > upper.x {
		r.x = upper.x
		r.width = lower.x - upper.x + 1
	} else {
		r.x = lower.x
		r.width = upper.x - lower.x + 1
	}

	if lower.y > upper.y {
		r.y = upper.y
		r.height = lower.y - upper.y + 1
	} else {
		r.y = lower.y
		r.height = upper.y - lower.y + 1
	}

	r.size = r.width * r.height
	r.lower = *lower
	r.upper = *upper

	return &r
}

func ComputeAnswer(grid *Grid) (answer int) {

	list := make([]*Rect, 0)
	for i := 0; i < len(grid.redTiles); i++ {
		for j := i + 1; j < len(grid.redTiles); j++ {
			r := MakeRect(grid.redTiles[i], grid.redTiles[j])
			list = append(list, r)
		}
	}

	slices.SortStableFunc(list, func(a, b *Rect) int {
		return b.size - a.size
	})
	// for i := range list {
	// 	log.Println(*list[i].lower, *list[i].upper, list[i])
	// }

	return list[0].size
}

func areLinesIntersecting(l1p1x, l1p1y, l1p2x, l1p2y, l2p1x, l2p1y, l2p2x, l2p2y float64) bool {
	q := (l1p1y-l2p1y)*(l2p2x-l2p1x) - (l1p1x-l2p1x)*(l2p2y-l2p1y)
	d := (l1p2x-l1p1x)*(l2p2y-l2p1y) - (l1p2y-l1p1y)*(l2p2x-l2p1x)
	if math.Abs(d) <= 0.1 {
		return false
	}
	r := q / d
	q = (l1p1y-l2p1y)*(l1p2x-l1p1x) - (l1p1x-l2p1x)*(l1p2y-l1p1y)
	s := q / d

	if r < 0 || r > 1 || s < 0 || s > 1 {
		return false
	}
	return true
}

func isPointPartOfTheLine(p *Point, p1x, p1y, p2x, p2y int) bool {

	return (p.x >= p1x && p.x <= p2x) && (p.y >= p1y && p.y <= p2y)
}

func IsPointInPolygon(p *Point, points []*Point) bool {
	j := len(points) - 1
	xHits := 0
	// yHits := 0
	isEdge := false
	// double ray cast
	// if point is on the same axis as edge, add buffer space for line intersection to work
	// return true if point is part of the edge
	for i := 0; i < len(points); i++ {
		p1x := points[i].x
		p1y := points[i].y
		p2x := points[j].x
		p2y := points[j].y
		if p1x > p2x {
			tmpx := p2x
			p2x, p1x = p1x, tmpx
		}
		if p1y > p2y {
			tmpy := p2y
			p2y, p1y = p1y, tmpy
		}
		isEdge = isPointPartOfTheLine(p, p1x, p1y, p2x, p2y)
		if isEdge {
			break
		}
		if p1x == p.x {
			if p1x >= p2x {
				p1x++
			} else {
				p1x--
			}
			// p2y++
		}
		if p2x == p.x {
			if p1x >= p2x {
				p2x++
			} else {
				p2x--
			}
		}
		xRayHit := areLinesIntersecting(float64(p.x), 0, float64(p.x), float64(p.y), float64(p1x), float64(p1y), float64(p2x), float64(p2y))
		if xRayHit {
			xHits++
		}
		// p1x = points[i].x
		// p2x = points[j].x

		// if p1y == p.y && p2y == p.y {
		// 	p1y--
		// 	p2y++
		// }
		// yRayHit := areLinesIntersecting(0, float64(p.y), float64(p.x), float64(p.y), float64(p1x), float64(p1y), float64(p2x), float64(p2y))
		// if yRayHit {
		// 	yHits++
		// }
		j = i

	}
	// log.Println("in1", inside)
	return isEdge || xHits > 0 && xHits%2 != 0
	// && yHits > 0 && yHits%2 != 0
}

// check if each corner of the rect is inside the polygon
func IsRectInsidePolygon(r *Rect, points []*Point) bool {
	p := Point{x: r.x, y: r.y}
	// top left
	if !IsPointInPolygon(&p, points) {
		return false
	}
	// top right
	p.x += r.width - 1
	if !IsPointInPolygon(&p, points) {
		return false
	}
	// bottom right
	p.y += r.height - 1
	if !IsPointInPolygon(&p, points) {
		return false
	}
	// bottom left
	p.x = r.x
	if !IsPointInPolygon(&p, points) {
		return false
	}

	// do deep check of every point
	// in case rectangle overlaps a hole or other shape
	for y := r.y; y < r.y+r.height-1; y += 100 {
		for x := r.x; x < r.x+r.width-1; x += 100 {
			p.x = x
			p.y = y
			inside := IsPointInPolygon(&p, points)
			if !inside {
				return false
			}
		}
	}
	return true
}

func ComputeAnswerPart2(grid *Grid) (answer int) {
	// f, _ := os.OpenFile("./out.txt", os.O_CREATE|os.O_TRUNC, 0666)
	// log.SetOutput(f)
	// defer f.Close()
	grid.max.x++
	grid.max.y++

	visualize := grid.max.x < 100

	if visualize {
		grid.items = make([][]byte, grid.max.y+1)
		for y := 0; y <= grid.max.y; y++ {
			grid.items[y] = make([]byte, grid.max.x+1)
			for x := 0; x <= grid.max.x; x++ {
				grid.items[y][x] = '.'
			}
		}
	}

	list := make([]*Rect, 0)
	for i := 0; i < len(grid.redTiles); i++ {
		if visualize {
			grid.items[grid.redTiles[i].y][grid.redTiles[i].x] = '#'
		}
		for j := i + 1; j < len(grid.redTiles); j++ {
			r := MakeRect(grid.redTiles[i], grid.redTiles[j])
			if IsRectInsidePolygon(r, grid.redTiles) {
				list = append(list, r)
			}
		}
	}
	if visualize {
		for y := 0; y <= grid.max.y; y++ {
			for x := 0; x <= grid.max.x; x++ {
				p := Point{x: x, y: y}
				if IsPointInPolygon(&p, grid.redTiles) {

					if grid.items[y][x] == '#' || grid.items[y][x] == '@' {
						grid.items[y][x] = '@'
					} else {
						grid.items[y][x] = 'x'
					}
				}
			}
		}
		PrintGrid(grid.items)
	}

	slices.SortStableFunc(list, func(a, b *Rect) int {
		return b.size - a.size
	})
	// for i := range list {
	// 	log.Println(list[i], list[i].size)
	// }

	if len(list) == 0 {
		return 0
	}
	return list[0].size
}

func main() {
	grid := ReadInput("./input.txt")
	answer := ComputeAnswer(&grid)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(&grid)
	log.Println("Answer part 2: ", answer)
}
