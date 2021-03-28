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

	scanner.Scan()
	numCases, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= numCases; i++ {
		scanner.Scan()
		list := strings.Split(scanner.Text(), " ")
		n, _ := strconv.ParseUint(list[0], 10, 64)
		k, _ := strconv.ParseUint(list[1], 10, 64)

		h, l := solve(n, k)
		writer.WriteString(fmt.Sprintf("Case #%d: %d %d\n", i, h, l))
		writer.Flush()
	}
}

func solve(n, k uint64) (uint64, uint64) {
	occupied := uint64(0)
	numDivisions := uint64(1)
	for occupied < k {
		occupied += numDivisions
		numDivisions *= 2
	}

	numDivisions /= 2
	occupied -= numDivisions

	lowNum := (n - occupied) / numDivisions
	numHighNums := (n - occupied) - lowNum*numDivisions

	remaining := k - occupied
	var toSplit uint64
	if remaining > numHighNums {
		toSplit = lowNum
	} else {
		toSplit = lowNum + 1
	}
	toSplit--
	if toSplit%2 == 0 {
		return toSplit / uint64(2), toSplit / uint64(2)
	}
	return toSplit/uint64(2) + uint64(1), toSplit / uint64(2)
}
