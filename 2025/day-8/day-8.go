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

type Vec3 struct {
	x int64
	y int64
	z int64
}

type TinyCircuit struct {
	a    Vec3
	b    Vec3
	dist float64
}

type Circuit struct {
	boxes []Vec3
}

// read each line in a file as an idrange
func ReadInput(path string) []Vec3 {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	boxes := make([]Vec3, 0)

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
			parts := strings.Split(s[0:endIndex], ",")
			numbers := make([]int64, len(parts))
			for i := range parts {
				num, e := strconv.ParseInt(parts[i], 10, 64)
				if e != nil {
					log.Panicln("Could not parse number, i:", i, ", s:", s)
				}
				numbers[i] = num
			}
			boxes = append(boxes, Vec3{x: numbers[0], y: numbers[1], z: numbers[2]})
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return boxes
}

func distance(v1 Vec3, v2 Vec3) float64 {
	dx := float64(v2.x - v1.x)
	dy := float64(v2.y - v1.y)
	dz := float64(v2.z - v1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func computePairsWithDistance(boxes []Vec3) []TinyCircuit {
	list := make([]TinyCircuit, 0, len(boxes))

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			sample := TinyCircuit{a: boxes[i], b: boxes[j]}
			sample.dist = distance(sample.a, sample.b)

			list = append(list, sample)
		}
	}

	slices.SortFunc(list, func(a TinyCircuit, b TinyCircuit) int {
		if a.dist < b.dist {
			return -1
		}
		return 0
	})

	return list
}

func ComputeAnswer(boxes []Vec3, maxConnections int) (answer int64) {
	list := computePairsWithDistance(boxes)
	// for i := range list[0 : maxConnections+20] {
	// 	log.Println("p", i, "", list[i])
	// }

	circuits := make([]*Circuit, len(boxes))
	for i := range boxes {
		circuits[i] = &Circuit{boxes: make([]Vec3, 1)}
		circuits[i].boxes[0] = boxes[i]
	}

	i := 0
	for {
		if i >= maxConnections {
			break
		}

		pair := list[i]
		i++
		indexOfA := slices.IndexFunc(circuits, func(e *Circuit) bool {
			return slices.ContainsFunc(e.boxes, func(a Vec3) bool {
				return pair.a == a
			})
		})
		indexOfB := slices.IndexFunc(circuits, func(e *Circuit) bool {
			return slices.ContainsFunc(e.boxes, func(a Vec3) bool {
				return pair.b == a
			})
		})
		// log.Println(i, pair, indexOfA, indexOfB)
		circuitA := circuits[indexOfA]
		circuitB := circuits[indexOfB]
		if !slices.Contains(circuitA.boxes, pair.b) {
			circuitA.boxes = append(circuitA.boxes, circuitB.boxes...)
			circuitB.boxes = make([]Vec3, 0)
		}
	}
	// cleanup empty circuits
	circuits = slices.DeleteFunc(circuits, func(e *Circuit) bool {
		return len(e.boxes) == 0
	})
	log.Println("connections made", i)

	slices.SortStableFunc(circuits, func(a, b *Circuit) int {
		return len(b.boxes) - len(a.boxes)
	})
	for i := range circuits {
		log.Println(i+1, circuits[i])
	}
	return int64(len(circuits[0].boxes) * len(circuits[1].boxes) * len(circuits[2].boxes))
}

func ComputeAnswerPart2(boxes []Vec3) (answer int64) {

	list := computePairsWithDistance(boxes)
	// for i := range list {
	// 	log.Println("p", i, "", list[i])
	// }

	circuits := make([]*Circuit, len(boxes))
	for i := range boxes {
		circuits[i] = &Circuit{boxes: make([]Vec3, 1)}
		circuits[i].boxes[0] = boxes[i]
	}

	i := 0
	var lastPair TinyCircuit
	for {
		if i >= len(list) {
			break
		}

		pair := list[i]
		i++
		indexOfA := slices.IndexFunc(circuits, func(e *Circuit) bool {
			return slices.ContainsFunc(e.boxes, func(a Vec3) bool {
				return pair.a == a
			})
		})
		indexOfB := slices.IndexFunc(circuits, func(e *Circuit) bool {
			return slices.ContainsFunc(e.boxes, func(a Vec3) bool {
				return pair.b == a
			})
		})
		// log.Println(i, pair, indexOfA, indexOfB)
		circuitA := circuits[indexOfA]
		circuitB := circuits[indexOfB]
		if !slices.Contains(circuitA.boxes, pair.b) {
			circuitA.boxes = append(circuitA.boxes, circuitB.boxes...)
			circuitB.boxes = make([]Vec3, 0)
		}
		// cleanup empty circuits
		circuits = slices.DeleteFunc(circuits, func(e *Circuit) bool {
			return len(e.boxes) == 0
		})

		if len(circuits) == 1 {
			lastPair = pair
			break
		}
	}
	log.Println("connections made", i)

	slices.SortStableFunc(circuits, func(a, b *Circuit) int {
		return len(b.boxes) - len(a.boxes)
	})
	for i := range circuits {
		log.Println(i+1, circuits[i])
	}
	return lastPair.a.x * lastPair.b.x
}

func main() {
	boxes := ReadInput("./input.txt")
	// answer := ComputeAnswer(boxes, 10) // for example input
	answer := ComputeAnswer(boxes, 1000) // for real input
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(boxes)
	log.Println("Answer part 2: ", answer)
}
