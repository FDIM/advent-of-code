package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Manifold struct {
	items [][]byte
}

// for part 2
type ManifoldPos struct {
	x, y int
}

// for part 2
type ManifoldPath struct {
	pos  *ManifoldPos
	prev *ManifoldPath
}

const SPOT_EMPTY = '.'
const SPOT_START = 'S'
const SPOT_SPLIT = '^'
const SPOT_BEAM = '|'

// read each line in a file as an battery
func ReadInput(path string) Manifold {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	items := make([][]byte, 0)
	manifold := Manifold{items: items}

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
			line := make([]byte, endIndex)
			for i := range s[0:endIndex] {
				line[i] = s[i]
			}
			// log.Println("s", s, cells)
			manifold.items = append(manifold.items, line)
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return manifold
}

func PrintManifold(items [][]byte) {
	for y := range items {
		b := strings.Builder{}
		for x := range items[y] {
			b.WriteRune(rune(items[y][x]))
		}
		log.Println(b.String())
	}
}

func IsSpotEmpty(manifold Manifold, x int, y int) bool {
	// out of bounds means empty
	if y < 0 || x < 0 || y >= len(manifold.items) || x >= len(manifold.items[y]) {
		return true
	}
	return manifold.items[y][x] == SPOT_EMPTY
}

func IsStart(manifold *Manifold, x int, y int) bool {
	return manifold.items[y][x] == SPOT_START
}

func IsSplitter(manifold *Manifold, x int, y int) bool {
	return manifold.items[y][x] == SPOT_SPLIT
}

func IsBeam(manifold *Manifold, x int, y int) bool {
	return manifold.items[y][x] == SPOT_BEAM
}

func ClearSpot(manifold *Manifold, x int, y int) {
	manifold.items[y][x] = SPOT_EMPTY
}

func CopyManifold(manifold *Manifold) *Manifold {
	copyManifold := Manifold{items: make([][]byte, len(manifold.items))}
	for y := range manifold.items {
		copyManifold.items[y] = make([]byte, len(manifold.items[y]))
		copy(copyManifold.items[y], manifold.items[y])
	}

	return &copyManifold
}

func ComputeAnswer(m *Manifold) (answer int64) {

	for y := 1; y < len(m.items); y++ {
		for x := range m.items[y] {
			// find start and place first beam
			if IsStart(m, x, y-1) {
				m.items[y][x] = SPOT_BEAM
				break
			}
			// find splitter, if it has beam above, place 2 beams on each side
			if IsSplitter(m, x, y) && IsBeam(m, x, y-1) {
				if !IsBeam(m, x-1, y) {
					m.items[y][x-1] = SPOT_BEAM
				}
				if !IsBeam(m, x+1, y) {
					m.items[y][x+1] = SPOT_BEAM
				}
				answer++
			}
			// pass through for beams
			if IsBeam(m, x, y-1) && !IsSplitter(m, x, y) {
				m.items[y][x] = SPOT_BEAM
			}
		}
	}

	PrintManifold(m.items)

	return answer
}

func ComputeAnswerPart2CopyMapSolution(entry *Manifold) (answer int) {
	// eventually this gets out of memory issue when y is 64+
	stack := make([]*Manifold, 1)
	stack[0] = entry
	for y := 1; y < len(entry.items); y++ {
		stackLen := len(stack)
		log.Println("Proccessing line: ", y, " stack size: ", stackLen)
		for mi := 0; mi < stackLen; mi++ {
			m := stack[mi]
			for x := range entry.items[y] {
				// find start and place first beam
				if IsStart(m, x, y-1) {
					m.items[y][x] = SPOT_BEAM
					break
				}
				// find splitter, if it has beam above, create another version of manifold
				// left goes on by itsel
				// right is copied before advancing, and goes by itself
				if IsSplitter(m, x, y) && IsBeam(m, x, y-1) {
					ml := CopyManifold(m)
					mr := CopyManifold(m)
					ml.items[y][x-1] = SPOT_BEAM
					mr.items[y][x+1] = SPOT_BEAM
					stack[mi] = ml
					stack = append(stack, mr)
					break
				}
				// pass through for beams
				if IsBeam(m, x, y-1) && !IsSplitter(m, x, y) {
					m.items[y][x] = SPOT_BEAM
				}
			}
		}
	}

	PrintManifold(stack[0].items)
	// for si := range stack {
	// 	PrintManifold(stack[si].items)
	// }
	return len(stack)
}

func UpdateManifoldState(m *Manifold, path *ManifoldPath, t byte) {
	current := path

	for {
		if current == nil || current.pos == nil {
			break
		}
		if m.items[current.pos.y][current.pos.x] == t {
			m.items[current.pos.y][current.pos.x] = 'E'
		} else {
			m.items[current.pos.y][current.pos.x] = t
		}
		current = current.prev
	}
}

func CreateManifoldPath(prev *ManifoldPath, pos *ManifoldPos) *ManifoldPath {
	p := ManifoldPath{}
	p.prev = prev
	p.pos = pos
	return &p
}

// much much better memory footprint but not good enough solution
// Reached 20gb usage when y is 86
// uses 1 map, but remembers change set for each variant using linked list
func ComputeAnswerPart2Pointers(entry *Manifold) (answer int) {
	stack := make([]*ManifoldPath, 1)
	stack[0] = &ManifoldPath{}
	m := CopyManifold(entry)

	for y := 1; y < len(entry.items); y++ {
		stackLen := len(stack)
		log.Println("Proccessing line: ", y, " stack size: ", stackLen)
		for pi := 0; pi < stackLen; pi++ {
			path := stack[pi]
			UpdateManifoldState(m, path, SPOT_BEAM)
			for x := range entry.items[y] {
				// find start and place first beam
				if IsStart(m, x, y-1) {
					path.pos = &ManifoldPos{x: x, y: y}
					break
				}
				// find splitter, if it has beam above, create another version of manifold
				// left goes on by itsel
				// right is copied before advancing, and goes by itself
				if IsSplitter(m, x, y) && IsBeam(m, x, y-1) {
					leftPath := CreateManifoldPath(path, &ManifoldPos{y: y, x: x - 1})
					rightPath := CreateManifoldPath(path, &ManifoldPos{y: y, x: x + 1})
					stack[pi] = leftPath
					stack = append(stack, rightPath)
					break
				}
				// pass through for beams
				if IsBeam(m, x, y-1) && !IsSplitter(m, x, y) {
					stack[pi] = CreateManifoldPath(path, &ManifoldPos{y: y, x: x})
				}
			}
			UpdateManifoldState(m, path, SPOT_EMPTY)
		}
	}

	// should print clean manifold
	PrintManifold(m.items)

	// for si := range stack {
	// 	UpdateManifoldState(m, stack[si], SPOT_BEAM)
	// 	PrintManifold(m.items)
	// 	UpdateManifoldState(m, stack[si], SPOT_EMPTY)
	// }
	return len(stack)
}

type Part2Result struct {
	answer int64
}
type BuildBeamAndUpdateManifoldArgs struct {
	x      int
	y      int
	offset int
}
type Part2Context struct {
	stack []BuildBeamAndUpdateManifoldArgs
	// used to cache splitter value that has been visited
	cache [][]int64
	path  *ManifoldPath
}

const GO_LEFT = -1
const GO_RIGHT = 1

func TraverseUpAndIncrementSplitterValue(m *Manifold, context *Part2Context, value int64) {
	// increment value
	p := context.path
	for {
		context.cache[p.pos.y][p.pos.x] += value
		p = p.prev
		if p == nil {
			break
		}
	}
	if len(context.stack) == 0 {
		return
	}
	// backtrack to the next right action that will run next
	p = context.path
	nextRightAction := context.stack[len(context.stack)-1]
	for {
		if (p.pos.x == nextRightAction.x && p.pos.y == nextRightAction.y) || p.prev == nil {
			break
		}
		context.path = p.prev
		p.prev = nil
		p = context.path

	}
}

func BuildBeamAndUpdateManifold(m *Manifold, sx int, sy int, offset int, context *Part2Context, res *Part2Result) {

	if m.items[sy][sx+offset] != SPOT_BEAM {
		m.items[sy][sx+offset] = SPOT_BEAM
	}

	for y := sy + 1; y < len(m.items); y++ {
		if IsSplitter(m, sx+offset, y) {
			// we have 2 beams on each splitter side, this means we have been there
			// and thus possible paths should be calculated already
			if m.items[y][sx+offset+GO_LEFT] == SPOT_BEAM && m.items[y][sx+offset+GO_RIGHT] == SPOT_BEAM {
				// log.Println("Reusing value for", y, sx+offset, context.cache[y][sx+offset])
				res.answer += context.cache[y][sx+offset]

				TraverseUpAndIncrementSplitterValue(m, context, context.cache[y][sx+offset])
			} else {
				splitterPath := &ManifoldPath{prev: context.path, pos: &ManifoldPos{x: sx + offset, y: y}}
				context.path = splitterPath
				context.stack = append(context.stack, BuildBeamAndUpdateManifoldArgs{x: sx + offset, y: y, offset: GO_RIGHT})
				BuildBeamAndUpdateManifold(m, sx+offset, y, GO_LEFT, context, res)
			}
			break
		} else {
			if m.items[y][sx+offset] != SPOT_BEAM {
				m.items[y][sx+offset] = SPOT_BEAM
			}

			if y == len(m.items)-1 {
				res.answer++
				TraverseUpAndIncrementSplitterValue(m, context, 1)
			}
		}
	}

}

func ComputeAnswerPart2(m *Manifold) (answer int64) {
	res := &Part2Result{answer: 0}
	sx := 0
	for x := range m.items[0] {
		// find start and place first beam
		if IsStart(m, x, 0) {
			sx = x
			break
		}
	}
	context := &Part2Context{stack: make([]BuildBeamAndUpdateManifoldArgs, 0), cache: make([][]int64, len(m.items))}
	for y := range m.items {
		context.cache[y] = make([]int64, len(m.items[0]))
	}
	context.path = &ManifoldPath{pos: &ManifoldPos{x: sx, y: 0}}
	BuildBeamAndUpdateManifold(m, sx, 1, 0, context, res)
	i := 0
	for {
		args := context.stack[len(context.stack)-1]
		context.stack = context.stack[0 : len(context.stack)-1]

		BuildBeamAndUpdateManifold(m, args.x, args.y, args.offset, context, res)

		i++
		// interrupt mid way to debug
		// if i >= 146 {
		// 	break
		// }
		if i > 1000000000 {
			PrintManifold(m.items)
			log.Println("last args: ", args, "current answer: ", res.answer)
			i = 0
		}
		if len(context.stack) == 0 {
			break
		}
	}
	UpdateManifoldState(m, context.path, '#')
	PrintManifold(m.items)
	// for c := range context.cache {
	// 	log.Println(context.cache[c])
	// }
	// answer = context.cache[0][sx] // same value as answer in the end
	answer = res.answer
	return answer
}

func main() {
	manifold := ReadInput("./input.txt")
	answer := ComputeAnswer(CopyManifold(&manifold))
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(CopyManifold(&manifold))
	log.Println("Answer part 2: ", answer)
}
