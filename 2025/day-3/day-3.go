package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type battery struct {
	cells []int
}

// read each line in a file as an battery
func ReadInput(path string) []battery {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	batteries := make([]battery, 0)

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
			cells := make([]int, endIndex) // exclude line ending

			for i := 0; i < endIndex; i++ {
				num, ee := strconv.Atoi(s[i : i+1])
				if ee != nil {
					log.Fatal("Failed to parse number, line: ", s, ee)
				}
				cells[i] = num
			}
			// log.Println("s", s, cells)
			batteries = append(batteries, battery{cells: cells})
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return batteries
}

func findMaxValueAndIndex(arr []int) (max int, index int) {
	max = 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			index = i
			max = arr[i]
		}
	}
	return max, index
}

func findCapacity(battery battery) (capacity int) {
	max, index := findMaxValueAndIndex(battery.cells[0 : len(battery.cells)-1])

	max2, _ := findMaxValueAndIndex(battery.cells[index+1:])
	capacity = max*10 + max2
	return capacity
}

func ComputeAnswer(batteries []battery) (answer int) {
	for i := range batteries {
		capacity := findCapacity(batteries[i])
		answer += capacity
		log.Println(batteries[i], capacity)
	}
	return answer
}

func findCapacityPart2(battery battery, count int) (capacity int) {
	res := make([]int, count)
	index := 0
	var sub []int
	log.Println("c", battery.cells)
	for i := range res {
		// easier to understand when cases are explicit
		// if i == 0 {
		// sub = battery.cells[0 : (len(battery.cells)+1)-count]
		// } else if i == count-1 {
		// sub = battery.cells[index:]
		// } else {
		sub = battery.cells[index : (len(battery.cells)+1)-(count-i)]
		// }
		max, indexOfMax := findMaxValueAndIndex(sub)
		res[i] = max
		index = indexOfMax + index + 1
		// log.Println("s", len(battery.cells), i, max, indexOfMax, index, sub)
		capacity += IntPow(10, count-1-i) * max
	}

	// capacity = max*10 + max2
	return capacity
}

func IntPow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}

func ComputeAnswerPart2(batteries []battery) (answer int) {
	for i := range batteries {
		capacity := findCapacityPart2(batteries[i], 12)
		answer += capacity
		log.Println(batteries[i], capacity)
	}
	return answer
}

func main() {
	batteries := ReadInput("./input.txt")
	answer := ComputeAnswer(batteries)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(batteries)
	log.Println("Answer part 2: ", answer)
}
