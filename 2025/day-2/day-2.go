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

type idrange struct {
	start uint64
	end   uint64
}

// read each line in a file as an idrange
func ReadInput(path string) []idrange {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	idranges := make([]idrange, 0)

	r := bufio.NewReader(f)
	for {
		s, e := r.ReadString(',')
		if len(s) > 0 {

			dashIndex := strings.Index(s, "-")
			if dashIndex == -1 {
				log.Println("Invalid range:", s)
				break
			}
			endIndex := len(s) - 1
			if e == io.EOF {
				endIndex++
			}
			start, startErr := strconv.ParseUint(s[0:dashIndex], 10, 64)
			end, endErr := strconv.ParseUint(s[dashIndex+1:endIndex], 10, 64)
			if startErr != nil || endErr != nil {
				log.Println("Invalid range, could not parse number:", s)
				break
			}
			idranges = append(idranges, idrange{start: start, end: end})
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return idranges
}

func IsValidId(id uint64) bool {
	s := strconv.FormatUint(id, 10)
	if len(s)%2 != 0 {
		return true
	}
	half := len(s) / 2
	// for i := 1; i <= half; i++ {
	subject := s[0:half]
	// log.Println("sub", half, subject, s)
	if strings.Contains(s[half:], subject) {
		return false
	}
	// }
	// log.Println(s, half)
	return true
}

func ComputeAnswer(idranges []idrange) (answer uint64) {
	for ri := range idranges {
		r := idranges[ri]
		for i := r.start; i <= r.end; i++ {
			valid := IsValidId(i)
			// log.Println("test:", i, ":", valid)
			if !valid {
				answer += i
			}
		}
		// log.Println(ri, r)
	}
	return answer
}

func IsValidIdForPart2(id uint64) (res bool) {
	res = true
	s := strconv.FormatUint(id, 10)
	half := len(s) / 2
	for i := 1; i <= half; i++ {
		subject := s[0:i]
		// log.Println("sub", i, subject, s)
		if strings.Contains(s[i:], subject) {
			repeating := isRepeating(s, subject)
			// log.Println("is repeating", s, subject, repeating)
			if repeating {
				res = false
				break
			}
		}
	}
	// log.Println(s, half)
	return res
}

func isRepeating(s string, subject string) bool {
	slen := len(subject)
	for i := slen; i < len(s); i += slen {
		if i+slen > len(s) || s[i:i+slen] != subject {
			return false
		}
	}
	return true
}

func ComputeAnswerPart2(idranges []idrange) (answer uint64) {

	for ri := range idranges {
		r := idranges[ri]
		for i := r.start; i <= r.end; i++ {
			valid := IsValidIdForPart2(i)
			// log.Println("test:", i, ":", valid)
			if !valid {
				answer += i
			}
		}
		// log.Println(ri, r)
	}
	return answer
}

func main() {
	idranges := ReadInput("./input.txt")
	answer := ComputeAnswer(idranges)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(idranges)
	log.Println("Answer part 2: ", answer)
}
