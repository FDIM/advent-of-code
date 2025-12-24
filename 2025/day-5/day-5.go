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

type NumRange struct {
	start int64
	end   int64
}

type Inventory struct {
	freshProductsRanges []NumRange
	stock               []int64
}

// read each line in a file as an battery
func ReadInput(path string) Inventory {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	inventory := Inventory{}
	inventory.freshProductsRanges = make([]NumRange, 0)
	inventory.stock = make([]int64, 0)
	isReadingStock := false
	r := bufio.NewReader(f)
	for {
		s, e := r.ReadString('\n')

		if s == "\n" {
			isReadingStock = true
			continue
		}
		if len(s) > 0 {
			endIndex := len(s) - 1
			// last line won't have \n if it's not there...
			// so don't skip last character
			if e == io.EOF {
				endIndex++
			}

			if isReadingStock {
				num, err := strconv.ParseInt(s[0:endIndex], 10, 64)
				if err != nil {
					log.Panicln("Could not parse number:", s)
				}
				inventory.stock = append(inventory.stock, num)
			} else {
				dashIndex := strings.Index(s, "-")
				if dashIndex == -1 {
					log.Panicln("Could not parse range, missing dash:", s)
				}
				start, startErr := strconv.ParseInt(s[0:dashIndex], 10, 64)
				end, endErr := strconv.ParseInt(s[dashIndex+1:endIndex], 10, 64)
				if startErr != nil || endErr != nil {
					log.Println("Invalid range, could not parse number:", s)
					break
				}

				inventory.freshProductsRanges = append(inventory.freshProductsRanges, NumRange{start: start, end: end})
			}
		}
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Fail to read line", e)
			break
		}
	}

	return inventory
}

func IsInRange(num int64, r NumRange) bool {
	return r.start <= num && r.end >= num
}

func IsProductFresh(id int64, rs []NumRange) (fresh bool) {
	for i := range rs {
		if IsInRange(id, rs[i]) {
			fresh = true
			break
		}
	}
	return fresh
}

func ComputeAnswer(inventory Inventory) (answer int64) {
	for i := range inventory.stock {
		if IsProductFresh(inventory.stock[i], inventory.freshProductsRanges) {
			answer++
		}
	}
	return answer
}

func dedupeRanges(ranges []NumRange) (list []NumRange, affected int) {
	affectedRanges := make([]NumRange, 0)
	list = make([]NumRange, 0)

	for i := range ranges {
		// skip if this range was incorporated somewhere else
		if slices.Contains(affectedRanges, ranges[i]) {
			continue
		}
		start := ranges[i].start
		end := ranges[i].end
		for ii := 0; ii < len(ranges); ii++ {
			// skip range under test
			if i == ii {
				continue
			}
			rr := ranges[ii]
			affected := false
			if IsInRange(start, rr) {
				start = rr.start
				affected = true
			}
			if IsInRange(end, rr) {
				end = rr.end
				affected = true
			}

			if affected {
				affectedRanges = append(affectedRanges, rr)
			}
		}
		list = append(list, NumRange{start: start, end: end})
	}
	affected = len(affectedRanges)
	return list, affected
}

func normalizeRanges(ranges []NumRange) []NumRange {
	var list []NumRange = ranges
	for {
		deduped, affected := dedupeRanges(list)
		log.Println(list, deduped, affected)
		list = deduped
		if affected == 0 {
			break
		}
	}

	return list
}

func ComputeAnswerPart2(inventory Inventory) (answer int64) {
	deduped := normalizeRanges(inventory.freshProductsRanges)
	log.Println("ranges ", inventory.freshProductsRanges)
	log.Println("deduped", deduped)
	for i := range deduped {
		r := deduped[i]
		answer += r.end - r.start + 1
	}
	return answer
}

func main() {
	inventory := ReadInput("./input.txt")
	answer := ComputeAnswer(inventory)
	log.Println("Answer part 1: ", answer)

	answer = ComputeAnswerPart2(inventory)
	log.Println("Answer part 2: ", answer)
}
