package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	if !scanner.Scan() {
		panic("number of test cases?")
	}
	numTestCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("couldn't read test case count")
	}
	for i := 0; i < numTestCases; i++ {
		calcEvacuation(i + 1)
	}
}

func calcEvacuation(testcase int) {
	if !scanner.Scan() {
		panic("couldn't read number of parties")
	}
	numParties, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("number of parties not a number")
	}
	if !scanner.Scan() {
		panic("couldn't read party numbers")
	}
	partyNums := map[string]int{}
	list := strings.Split(scanner.Text(), " ")
	for i := 0; i < numParties; i++ {
		num, err := strconv.Atoi(list[i])
		if err != nil {
			panic("couldn't read party number")
		}
		partyNums[fmt.Sprintf("%c", byte('A'+i))] = num
	}

	output := []string{}
	for total(partyNums) > 0 {
		output = append(output, getNextEvac(partyNums))
	}
	writer.WriteString(fmt.Sprintf("Case #%d: %s\n", testcase, strings.Join(output, " ")))
	writer.Flush()
}

func getNextEvac(partyNums map[string]int) string {
	p1, _ := getMax(partyNums)
	partyNums[p1]--
	p2, _ := getMax(partyNums)
	partyNums[p2]--
	if noMajority(partyNums) {
		return fmt.Sprintf("%s%s", p1, p2)
	}
	partyNums[p2]++
	return p1
}

func getMax(partyNums map[string]int) (string, int) {
	var maxk string
	maxv := 0
	for k, v := range partyNums {
		if v > maxv {
			maxv = v
			maxk = k
		}
	}
	return maxk, maxv
}

func total(partyNums map[string]int) int {
	total := 0
	for _, v := range partyNums {
		total += v
	}
	return total
}

func noMajority(partyNums map[string]int) bool {
	total := total(partyNums)
	_, max := getMax(partyNums)
	return 2*max <= total
}
