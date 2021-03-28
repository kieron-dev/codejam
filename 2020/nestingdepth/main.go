package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numCases int
	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		panic(err)
	}

	for ncase := 0; ncase < numCases; ncase++ {
		var s string
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			panic(err)
		}

		var nums []int
		for _, c := range s {
			nums = append(nums, int(byte(c)-'0'))
		}

		res := calc(nums)
		fmt.Printf("Case #%d: %s\n", ncase+1, res)
	}
}

func calc(nums []int) string {
	level := 0
	out := ""

	for _, n := range nums {
		for level < n {
			out += "("
			level++
		}
		for level > n {
			out += ")"
			level--
		}
		out += strconv.Itoa(n)
	}

	for level > 0 {
		out += ")"
		level--
	}

	return out
}
