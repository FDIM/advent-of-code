package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type HomeworkProblem struct {
	numbers   []int64
	operation byte
}

type Homework struct {
	lines []string
}

const OP_ADD = '+'
const OP_MULTIPLY = '*'

// read each line in a file
func ReadInput(path string) Homework {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	homework := Homework{}
	homework.lines = make([]string, 0)
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
			homework.lines = append(homework.lines, s[0:endIndex])
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return homework
}

func ReadProblems(homework Homework) []*HomeworkProblem {
	problems := make([]*HomeworkProblem, 0)
	for si := range homework.lines {
		numReader := bufio.NewReader(strings.NewReader(homework.lines[si]))
		i := 0
		for {
			numStr, ee := numReader.ReadString(' ')

			if len(numStr) > 0 && numStr != " " {
				var problem *HomeworkProblem

				if i >= len(problems) {
					problem = &HomeworkProblem{}
					problem.numbers = make([]int64, 0)
					problems = append(problems, problem)
				} else {
					problem = problems[i]
				}

				if numStr[0] == OP_ADD || numStr[0] == OP_MULTIPLY {
					problem.operation = numStr[0]
				} else {
					endNumIndex := len(numStr) - 1
					if ee == io.EOF {
						endNumIndex++
					}
					// ^ only needed for part 2
					num, err := strconv.ParseInt(numStr[0:endNumIndex], 10, 64)
					if err != nil {
						log.Panicln("Could not parse number:", numStr)
					}
					problem.numbers = append(problem.numbers, num)
				}
				i++
			}
			if ee == io.EOF {
				break
			}
		}
	}
	return problems
}

func SolveProblem(problem *HomeworkProblem) (res int64) {
	res = problem.numbers[0]
	for i := 1; i < len(problem.numbers); i++ {
		num := problem.numbers[i]
		switch problem.operation {
		case OP_ADD:
			res += num
		case OP_MULTIPLY:
			res *= num
		default:
			panic("Unknown operation")
		}
	}
	return res
}

func ComputeAnswer(homework Homework) (answer int64) {
	problems := ReadProblems(homework)
	for i := range problems {
		p := *problems[i]
		res := SolveProblem(&p)
		answer += res
		log.Println(p, res)
	}
	return answer
}

func ReadProblemsPart2(homework Homework) []*HomeworkProblem {
	problems := make([]*HomeworkProblem, 0)
	to := 0
	from := 0
	line := homework.lines[0]
	// when a whitespace is found, check if all lines have a whitespace there
	// if yes, parse the problem
	for i := 0; i <= len(line); i++ {
		// intentinally going out of bounds!
		if i < len(line) && line[i] != ' ' {
			continue
		}
		allLinesHasSpace := true
		if i < len(line) {
			for si := range homework.lines[1:len(homework.lines)] {
				if homework.lines[si][i] != ' ' {
					allLinesHasSpace = false
					break
				}
			}
		}
		if allLinesHasSpace {
			to = from
			from = i
			problems = append(problems, ParseProblemPart2(homework, from, to))
		}
	}
	return problems
}

// reading from right to left, so from is bigger than to
// range is relative to the gap!
func ParseProblemPart2(homework Homework, from int, to int) *HomeworkProblem {
	if to != 0 {
		to++
	}
	from -= 1
	problem := HomeworkProblem{numbers: make([]int64, 0)}
	for i := from; i >= to; i-- {
		str := ""
		for l := range homework.lines[0 : len(homework.lines)-1] {
			char := homework.lines[l][i : i+1]
			if char != " " {
				str += char
			}
		}
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Panicln("Could not parse number:", str, from, to)
		}
		problem.numbers = append(problem.numbers, num)
		problem.operation = homework.lines[len(homework.lines)-1][to : to+1][0]
	}
	// log.Println("s", problem)
	return &problem
}

func ComputeAnswerPart2(homework Homework) (answer int64) {
	problems := ReadProblemsPart2(homework)
	for i := range problems {
		p := *problems[i]
		res := SolveProblem(&p)
		answer += res
		log.Println(p, res)
	}
	return answer
}

func main() {
	homework := ReadInput("./input.txt")
	answer := ComputeAnswer(homework)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(homework)
	log.Println("Answer part 2: ", answer)
}
