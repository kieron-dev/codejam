package main

import (
	"fmt"
	"log"
)

func main() {
	var numCases int

	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for n := 1; n <= numCases; n++ {
		var x, y int
		var s string
		_, err := fmt.Scanf("%d %d %s", &x, &y, &s)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Case #%d: %d\n", n, solve(x, y, s))
	}
}

func solve(x, y int, input string) int {
	costs := map[[2]byte]int{
		{'C', 'C'}: 0,
		{'J', 'J'}: 0,
		{'C', 'J'}: x,
		{'J', 'C'}: y,
	}

	if input[0] == '?' {
		a := solve(x, y, "C"+input[1:])
		b := solve(x, y, "J"+input[1:])
		if a < b {
			return a
		}
		return b
	}

	prev := input[0]
	cost := 0

	for i := 1; i < len(input); i++ {
		cur := input[i]
		if cur == '?' {
			if costs[[2]byte{prev, byte('C')}] < costs[[2]byte{prev, byte('J')}] {
				cur = 'C'
			} else {
				cur = 'J'
			}
		}
		cost += costs[[2]byte{prev, cur}]
		prev = cur
	}

	return cost
}
