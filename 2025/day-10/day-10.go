package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	state         []string
	target        []string
	buttons       []*Button
	joltage       []int
	joltageTarget []int
	ready         bool
}

type Button struct {
	indexes []int
}

const STATE_OFF = "."
const STATE_ON = "#"

// read each line in a file as an idrange
func ReadInput(path string) []*Machine {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	machines := make([]*Machine, 0)

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

			machines = append(machines, MakeMachine(s[0:endIndex]))
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return machines
}

func MakeMachine(s string) *Machine {
	m := Machine{}
	m.buttons = make([]*Button, 0)
	parts := strings.Split(s, " ")

	for pi := range parts {
		p := parts[pi]

		switch p[0] {
		case '[':
			m.target = strings.Split(p[1:len(p)-1], "")
			m.state = make([]string, len(m.target))
			for i := range m.target {
				m.state[i] = STATE_OFF
			}
		case '(':
			m.buttons = append(m.buttons, &Button{indexes: ParseNumbers(p)})
		case '{':
			m.joltageTarget = ParseNumbers(p)
			m.joltage = make([]int, len(m.joltageTarget))
		}
	}

	return &m
}

func ParseNumbers(s string) []int {
	parts := strings.Split(s[1:len(s)-1], ",")
	numbers := make([]int, len(parts))
	for i := range parts {
		num, e := strconv.Atoi(parts[i])
		if e != nil {
			log.Panicln("Could not parse number, i:", i, ", s:", s)
		}
		numbers[i] = num
	}

	return numbers
}

func PrintMachine(m *Machine) {
	log.Println(m.target, ":", m.state, m.ready, stringifyButtons(m.buttons), m.joltage, m.joltageTarget)
}

func stringifyButtons(buttons []*Button) string {
	if len(buttons) == 0 {
		return "()"
	}
	str := ""
	for a := range buttons {
		indexes := make([]string, len(buttons[a].indexes))
		for i := range buttons[a].indexes {
			indexes[i] = strconv.Itoa(buttons[a].indexes[i])
		}
		str += "(" + strings.Join(indexes, ",") + ") "
	}

	return str[0 : len(str)-1]
}

func CopyMachine(src *Machine) *Machine {
	m := Machine{}
	m.buttons = src.buttons
	m.ready = false
	m.joltageTarget = src.joltageTarget
	m.target = src.target
	m.state = make([]string, len(src.target))
	for i := range src.target {
		m.state[i] = STATE_OFF
	}
	m.joltage = make([]int, len(src.joltageTarget))
	return &m
}

type MachineButtonPressFn func(m *Machine, b *Button)

// part 1
func PressMachineButtonToUpdateState(m *Machine, b *Button) {
	for i := range b.indexes {
		index := b.indexes[i]
		if m.state[index] == STATE_ON {
			m.state[index] = STATE_OFF
		} else {
			m.state[index] = STATE_ON
		}
	}

	m.ready = IsStateMatching(m)
}

func IsStateMatching(m *Machine) bool {
	for i := range m.target {
		if m.target[i] != m.state[i] {
			return false
		}
	}
	return true
}

// part 2
func PressMachineButtonToUpdateJoltage(m *Machine, b *Button) {
	for i := range b.indexes {
		index := b.indexes[i]
		m.joltage[index]++
	}
	m.ready = IsJoltageMatching(m)
}

func IsJoltageMatching(m *Machine) bool {
	for i := range m.joltage {
		if m.joltage[i] != m.joltageTarget[i] {
			return false
		}
	}
	return true
}

func IsJoltageTooBig(m *Machine) bool {
	for i := range m.joltage {
		if m.joltage[i] > m.joltageTarget[i] {
			return true
		}
	}
	return false
}

func ComputeNextRelatedButtons(m *Machine, context *SimulationContext, previousActions []*Button) [][]*Button {
	list := make([][]*Button, 0)
	if !context.part2 {
		// worked for part 1
		for b := range m.buttons {
			button := m.buttons[b]
			if len(previousActions) == 0 || previousActions[len(previousActions)-1] != button {
				list = append(list, make([]*Button, 1))
				list[len(list)-1][0] = m.buttons[b]
			}
		}
	} else if !IsJoltageTooBig(m) {

		// TODO: figure logic for part 2
		//
		// for i := range m.joltage {
		// 	// if m.joltage[i] < m.joltageTarget[i] {
		// 	for s := m.joltage[i] + 1; s <= m.joltageTarget[i]; s++ {
		// 		for si := 0; si < s; si++ {
		// 			actionsSet := make([]*Button, 0)
		// 			for b := range m.buttons {
		// 				button := m.buttons[b]
		// 				if slices.Contains(button.indexes, i) {
		// 					actionsSet = append(actionsSet, button)
		// 				}
		// 			}
		// 			// log.Println(s, stringifyButtons(actionsSet))
		// 			list = append(list, actionsSet)
		// 		}
		// 		// }
		// 		// log.Println(s)
		// 	}
		// 	// log.Panicln("aaa", list)
		// }
	}
	return list
}

type SimulationContext struct {
	maxIterations int
	list          [][]*Button
	buttonPressFn MachineButtonPressFn
	part2         bool
}

func SimulateButtonPresses(m *Machine, context *SimulationContext, previousActions []*Button) {

	if previousActions == nil {
		previousActions = make([]*Button, 0)
	}

	nm := CopyMachine(m)
	for a := range previousActions {
		context.buttonPressFn(nm, previousActions[a])
	}

	// log.Println("nm", stringifyButtons(previousActions))
	// PrintMachine(nm)

	if nm.ready {
		context.list = append(context.list, previousActions)
	} else if len(previousActions) < context.maxIterations {
		actions := ComputeNextRelatedButtons(nm, context, previousActions)
		for a := range actions {
			for s := range actions[a] {
				list := make([]*Button, 0)
				button := actions[a][s]
				list = append(list, previousActions...)
				list = append(list, button)
				SimulateButtonPresses(m, context, list)
			}
		}
	}
}

func SimulateAsync(m *Machine, i int, fn MachineButtonPressFn, part2 bool, ch chan int) {
	maxAllowedIterations := 100
	PrintMachine(m)
	context := &SimulationContext{list: make([][]*Button, 0)}
	context.maxIterations = 1
	context.buttonPressFn = fn
	context.part2 = part2
	for {
		SimulateButtonPresses(m, context, nil)

		if len(context.list) > 0 {
			break
		} else if context.maxIterations > maxAllowedIterations {
			log.Println("max iteration reached for ", i)
			break
		} else {
			context.maxIterations *= 2
		}
	}
	if len(context.list) == 0 {
		ch <- -1
		return
	}
	slices.SortStableFunc(context.list, func(a, b []*Button) int {
		return len(a) - len(b)
	})
	// for j := range res.list {
	j := 0 // take first, as it will have smallest number of actions
	nm := CopyMachine(m)
	for a := range context.list[j] {
		context.buttonPressFn(nm, context.list[j][a])
	}
	log.Println(nm.ready, len(context.list[j]), stringifyButtons(context.list[j]))
	PrintMachine(nm)
	ch <- len(context.list[j])
}

func ComputeAnswer(machines []*Machine, fn MachineButtonPressFn, part2 bool) (answer int) {
	channels := make([]chan int, 0)
	for i := range machines {
		m := machines[i]
		ch := make(chan int)
		channels = append(channels, ch)
		go SimulateAsync(m, i, fn, part2, ch)

		// debug code for the first machine
		// break
	}

	for i := range channels {
		min := <-channels[i]
		log.Println("Progress:", i, "/", len(channels))
		if min == -1 {
			log.Println("Failed to find count for ", i)
			PrintMachine(machines[i])
		}
		answer += min
	}

	return answer
}

func ComputeAnswerPart2(machines []*Machine) (answer int) {
	answer = ComputeAnswer(machines, PressMachineButtonToUpdateJoltage, true)
	return answer
}

func main() {
	machines := ReadInput("./input.example.txt")
	answer := ComputeAnswer(machines, PressMachineButtonToUpdateState, false)
	log.Println("Answer part 1: ", answer)

	answer2 := ComputeAnswerPart2(machines)
	log.Println("Answer part 2: ", answer2)
}
