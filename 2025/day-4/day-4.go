package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Diagram struct {
	items           [][]int
	accessibleItems [][]int
}

const SPOT_EMPTY = int('.')
const SPOT_ROLL = int('@')
const SPOT_MARK = int('X')

// read each line in a file as an battery
func ReadInput(path string) Diagram {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	items := make([][]int, 0)
	diagram := Diagram{items: items}

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
			line := make([]int, endIndex)
			for i := range s[0:endIndex] {
				line[i] = int(s[i])
			}
			// log.Println("s", s, cells)
			diagram.items = append(diagram.items, line)
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return diagram
}

func PrintDiagram(items [][]int) {
	for y := range items {
		b := strings.Builder{}
		for x := range items[y] {
			b.WriteRune(rune(items[y][x]))
		}
		log.Println(b.String())
	}
}

func IsSpotEmpty(diagram Diagram, x int, y int) bool {
	// out of bounds means empty
	if y < 0 || x < 0 || y >= len(diagram.items) || x >= len(diagram.items[y]) {
		return true
	}
	return diagram.items[y][x] == SPOT_EMPTY
}

func IsRoll(diagram Diagram, x int, y int) bool {
	return diagram.items[y][x] == SPOT_ROLL
}

func CountAdjacentRolls(diagram Diagram, x int, y int) (count int) {
	// check top
	if !IsSpotEmpty(diagram, x, y-1) {
		count++
	}
	// check bottom
	if !IsSpotEmpty(diagram, x, y+1) {
		count++
	}
	// check each left side
	if !IsSpotEmpty(diagram, x-1, y-1) {
		count++
	}
	if !IsSpotEmpty(diagram, x-1, y) {
		count++
	}
	if !IsSpotEmpty(diagram, x-1, y+1) {
		count++
	}

	// check each right side
	if !IsSpotEmpty(diagram, x+1, y-1) {
		count++
	}
	if !IsSpotEmpty(diagram, x+1, y) {
		count++
	}
	if !IsSpotEmpty(diagram, x+1, y+1) {
		count++
	}
	return count
}

func ComputeAnswer(diagram Diagram) (answer int) {
	diagram.accessibleItems = make([][]int, len(diagram.items))
	for y := range diagram.items {
		diagram.accessibleItems[y] = make([]int, len(diagram.items[y]))
		copy(diagram.accessibleItems[y], diagram.items[y])
		for x := range diagram.items[y] {
			if IsRoll(diagram, x, y) {
				count := CountAdjacentRolls(diagram, x, y)
				if count < 4 {
					diagram.accessibleItems[y][x] = SPOT_MARK
					answer += 1
				}
			}
		}
	}

	PrintDiagram(diagram.accessibleItems)

	return answer
}

func CountAndRemoveRolls(diagram Diagram) (rolls int) {
	for y := range diagram.items {
		for x := range diagram.items[y] {
			if IsRoll(diagram, x, y) {
				count := CountAdjacentRolls(diagram, x, y)
				if count < 4 {
					// apparently defering is not important
					defer ClearSpot(diagram, x, y)
					rolls += 1
				}
			}
		}
	}

	return rolls
}

func ClearSpot(diagram Diagram, x int, y int) {
	diagram.items[y][x] = SPOT_EMPTY
}

func CopyDiagram(diagram Diagram) Diagram {
	copiedDiagram := Diagram{items: make([][]int, len(diagram.items))}
	for y := range diagram.items {
		copiedDiagram.items[y] = make([]int, len(diagram.items[y]))
		copy(copiedDiagram.items[y], diagram.items[y])
	}

	return copiedDiagram
}

func ComputeAnswerPart2(diagram Diagram) (answer int) {
	copy := CopyDiagram(diagram)
	for {
		PrintDiagram(copy.items)
		count := CountAndRemoveRolls(copy)
		log.Println("removed", count)
		answer += count
		if count == 0 {
			break
		}
	}
	return answer
}

func main() {
	diagram := ReadInput("./input.txt")
	answer := ComputeAnswer(diagram)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(diagram)
	log.Println("Answer part 2: ", answer)
}
