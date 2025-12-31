package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

const PATH_START = "you"
const PATH_END = "out"
const PATH_START_PART2 = "svr"
const PATH_FAULTY_1 = "dac"
const PATH_FAULTY_2 = "fft"

type Device struct {
	id      string
	outputs []string
}

// read each line in a file as an idrange
func ReadInput(path string) []*Device {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	devices := make([]*Device, 0)

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

			devices = append(devices, MakeDevice(s[0:endIndex]))
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return devices
}

func MakeDevice(s string) *Device {
	m := Device{}
	colonIndex := strings.Index(s, ": ")
	m.id = s[0:colonIndex]
	m.outputs = strings.Split(s[colonIndex+2:], " ")
	return &m
}

type SearchContext struct {
	paths        [][]string
	visitedCount int64
	visited      map[string]int
}

func TraverseTree(devices []*Device, start string, context *SearchContext, requiredSequence []string) {
	TraverseTreeImpl(devices, start, context, requiredSequence, make([]string, 0), start, strconv.Itoa(rand.Int()))
}

func TraverseTreeImpl(devices []*Device, start string, context *SearchContext, requiredSequence []string, path []string, target string, prefix string) {

	// pathKey := prefix + ":" + target
	// count, isKnownPath := context.visited[pathKey]
	// // nothing to do if we have tried this path already and found nothing
	// if count == 0 && isKnownPath {
	// 	return
	// }
	path = append(path, target)

	if target == requiredSequence[0] {
		// log.Println("Found path to", requiredSequence[0], path, len(requiredSequence))
		requiredSequence = requiredSequence[1:]

		if len(requiredSequence) == 0 {
			context.paths = append(context.paths, path)
			return
		}
		// prefix = strconv.Itoa(rand.Int())
	}

	index := slices.IndexFunc(devices, MakeDeviceComparatorById(target))
	// we probably found the end too soon
	if index == -1 {
		return
	}

	d := devices[index]

	if context.visitedCount%100000 == 0 {
		log.Println(target, path, requiredSequence, context.visitedCount, len(context.paths))
	}

	context.visitedCount++

	for i := 0; i < len(d.outputs); i++ {
		// if i == 0 {
		// pathsCount := len(context.paths)
		TraverseTreeImpl(devices, start, context, requiredSequence, path, d.outputs[i], prefix)
		// context.visited[pathKey] = len(context.paths) - pathsCount
		// } else {
		// 	defer func(ii int) {
		// 		pathsCount := len(context.paths)
		// 		TraverseTreeImpl(devices, start, context, requiredSequence, path, d.outputs[i], prefix)
		// 		context.visited[pathKey] = len(context.paths) - pathsCount
		// 	}(i)
		// }
	}
}

func MakeDeviceComparatorById(id string) func(d *Device) bool {
	return func(d *Device) bool {
		return d.id == id
	}
}

func ComputeAnswer(devices []*Device) (answer int) {

	context := SearchContext{paths: make([][]string, 0), visited: make(map[string]int)}

	TraverseTree(devices, PATH_START, &context, []string{PATH_END})
	// log.Println(context.visited)
	// log.Println(context.paths)
	answer = len(context.paths)
	return answer
}

func ComputeAnswerPart2(devices []*Device) (answer int) {
	context := SearchContext{paths: make([][]string, 0), visited: make(map[string]int)}

	// TraverseTree(devices, PATH_START_PART2, &context, []string{PATH_END})
	TraverseTree(devices, PATH_START_PART2, &context, []string{PATH_FAULTY_1, PATH_FAULTY_2, PATH_END})
	TraverseTree(devices, PATH_START_PART2, &context, []string{PATH_FAULTY_2, PATH_FAULTY_1, PATH_END})

	// log.Println(context.visited)
	// log.Println(context.paths)
	answer = len(context.paths)
	return answer
}

func main() {
	devices := ReadInput("./input.txt")
	// devices := ReadInput("./input.example.txt")
	answer := ComputeAnswer(devices)
	log.Println("Answer part 1: ", answer)

	devices = ReadInput("./input.example.part2.txt")
	answer2 := ComputeAnswerPart2(devices)
	log.Println("Answer part 2: ", answer2)
}
