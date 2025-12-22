package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type action struct {
	right bool
	count int
}

// Hello returns a greeting for the named person.
func ReadInput(path string) []action {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	actions := make([]action, 0)
	r := bufio.NewReader(f)
	s, e := r.ReadString('\n')
	for e == nil {

		right := s[0] == 'R'
		count, ee := strconv.Atoi(s[1 : len(s)-1])

		if ee != nil {
			log.Fatal("Failed to parse number, line: ", s, ee)
		}

		actions = append(actions, action{count: count, right: right})
		s, e = r.ReadString('\n')
	}
	return actions
}

func ComputeAnswer(actions []action) (answer int) {
	dial := 50
	for _, a := range actions {
		count := a.count
		if count > 99 {
			count %= 100
		}
		if a.right {
			dial += count
		} else {
			dial -= count
		}
		if dial < 0 {
			dial += 100
		}
		if dial > 99 {
			dial -= 100
		}
		// log.Println("a", a.right, dial)
		if dial == 0 {
			answer++
		}
	}
	log.Println(dial)
	return answer
}

func ComputeAnswerPart2(actions []action) (answer int) {
	dial := 50
	for _, a := range actions {
		for i := 0; i < a.count; i++ {
			if a.right {
				dial++
			} else {
				dial--
			}
			if dial < 0 {
				dial += 100
				// answer++
			}
			if dial > 99 {
				dial -= 100
				// answer++
			}
			if dial == 0 {
				answer++
			}
			// log.Println("a", a, dial, answer)
		}
	}
	log.Println(dial)
	return answer
}

func main() {
	actions := ReadInput("./input.txt")
	answer := ComputeAnswer(actions)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(actions)
	log.Println("Answer part 2: ", answer)
}
