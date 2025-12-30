package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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
	paths [][]string
	stack []SearchContextStack
	// unsure if needed
	visited int64
}

type SearchContextStack struct {
	index            int
	startIndex       int
	requiredSequence []string
	path             []string
}

func TraverseTree(devices []*Device, index int, context *SearchContext, requiredSequence []string) {
	TraverseTreeImpl(devices, index, context, requiredSequence, make([]string, 0), index)
}

func TraverseTreeImpl(devices []*Device, startIndex int, context *SearchContext, requiredSequence []string, path []string, index int) {

	d := devices[index]

	// unsure if needed
	// if context.visited != nil && len(path) > 0 {
	// 	str := strings.Join(requiredSequence, ",") + "|" + strings.Join(path, ",") + d.id
	// 	if slices.Contains(context.visited, str) {
	// 		log.Println("duplicate", str)
	// 		return
	// 	}
	// 	context.visited = append(context.visited, str)
	// }

	if context.visited%10000 == 0 {
		log.Println(d.id, path, requiredSequence, context.visited, len(context.stack))
	}
	context.visited++

	for i := 0; i < len(d.outputs); i++ {
		outId := d.outputs[i]

		if outId == requiredSequence[0] {
			path = append(path, requiredSequence[0])
			log.Println("Found path to", requiredSequence[0], path, len(requiredSequence))
			requiredSequence = requiredSequence[1:]

			if len(requiredSequence) == 0 {
				context.paths = append(context.paths, path)
				break
			}
		}
		// we found end too soon and it does not follow the required sequence
		if outId == PATH_END {
			break
		}
		nextIndex := slices.IndexFunc(devices, MakeDeviceComparatorById(outId))
		// log.Println(outId, index, nextIndex)
		npath := make([]string, len(path))
		copy(npath, path)
		npath = append(npath, d.id)
		// if i == 0 {
		// 	TraverseTreeImpl(devices, startIndex, context, requiredSequence, path, nextIndex)
		// } else {
			context.stack = append(context.stack, SearchContextStack{startIndex: startIndex, index: nextIndex, requiredSequence: requiredSequence, path: npath})
		// }
	}
}

func RunSearchStack(devices []*Device, context *SearchContext) {
	for {
		if len(context.stack) == 0 {
			break
		}
		s := context.stack[0]
		TraverseTreeImpl(devices, s.startIndex, context, s.requiredSequence, s.path, s.index)
		context.stack = context.stack[1:]
	}
}

func MakeDeviceComparatorById(id string) func(d *Device) bool {
	return func(d *Device) bool {
		return d.id == id
	}
}

func ComputeAnswer(devices []*Device) (answer int) {
	// data only goes one way, so skip everything until "you" device
	start := slices.IndexFunc(devices, MakeDeviceComparatorById(PATH_START))

	context := SearchContext{paths: make([][]string, 0)}

	TraverseTree(devices, start, &context, []string{PATH_END})
	RunSearchStack(devices, &context)
	// log.Println(context.paths)
	answer = len(context.paths)
	return answer
}

func ComputeAnswerPart2(devices []*Device) (answer int) {
	// data only goes one way, so skip everything until "you" device
	start := slices.IndexFunc(devices, MakeDeviceComparatorById(PATH_START_PART2))

	context := SearchContext{paths: make([][]string, 0)}

	TraverseTree(devices, start, &context, []string{PATH_FAULTY_1, PATH_FAULTY_2, PATH_END})
	RunSearchStack(devices, &context)
	TraverseTree(devices, start, &context, []string{PATH_FAULTY_2, PATH_FAULTY_1, PATH_END})
	RunSearchStack(devices, &context)

	// for p := range context.visited {
	// 	log.Println("v", context.visited[p])
	// }
	log.Println(context.paths)
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
