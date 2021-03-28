package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var numCases int

	_, err := fmt.Scanf("%d\n", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numCases; i++ {
		var input string

		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			log.Fatal(err)
		}

		a, b := solve(input)
		fmt.Printf("Case #%d: %s %s\n", i+1, a, b)
	}
}

func solve(nStr string) (a, b string) {
	for i := len(nStr) - 1; i >= 0; i-- {
		char := nStr[i : i+1]
		digit, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		if digit == 4 {
			a = "3" + a
			b = "1" + b
		} else {
			a = char + a
			b = "0" + b
		}
	}

	i := 0
	for b[i:i+1] == "0" {
		i++
	}
	return a, b[i:]
}
