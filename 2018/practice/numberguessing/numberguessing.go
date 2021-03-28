package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("didn't get expected number of tests")
	}
	numTests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "numTests = %+v\n", numTests)

	for i := 0; i < numTests; i++ {
		if !guessNums(scanner) {
			return
		}
	}
}

func guessNums(scanner *bufio.Scanner) bool {
	if !scanner.Scan() {
		panic("didn't get expected limits")
	}

	parts := strings.Split(scanner.Text(), " ")
	if len(parts) != 2 {
		panic("limits malformed")
	}
	lowBound, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	highBound, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "lowBound = %+v, highBound = %+v\n", lowBound, highBound)

	if !scanner.Scan() {
		panic("didn't get expected attempt count")
	}
	attempts, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Attempt count malformed")
	}
	fmt.Fprintf(os.Stderr, "attempts = %+v\n", attempts)

	return interactiveGuesses(scanner, lowBound, highBound, attempts)
}

func interactiveGuesses(scanner *bufio.Scanner, low, high, lim int) bool {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < lim; i++ {
		diff := high - low
		var next int
		if diff == 1 {
			next = low + 1
		} else {
			pow := 1
			for pow < diff {
				pow *= 2
			}
			pow /= 2
			next = low + pow
		}
		fmt.Fprintln(os.Stderr, low, high, next)
		writer.WriteString(fmt.Sprintf("%d\n", next))
		writer.Flush()
		if !scanner.Scan() {
			panic("didn't get response")
		}
		resp := scanner.Text()
		switch resp {
		case "CORRECT":
			fmt.Fprintf(os.Stderr, "Correct answer %d in %d attempts out of %d\n", next, i+1, lim)
			return true
		case "WRONG_ANSWER":
			return false
		case "TOO_SMALL":
			low = next
		case "TOO_BIG":
			high = next - 1
		default:
			fmt.Fprintf(os.Stderr, "Unexpected response: %s\n", resp)
			return false
		}
	}
	return false
}
