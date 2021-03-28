package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var numCases int

	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for n := 1; n <= numCases; n++ {
		var numNums, cost int
		_, err := fmt.Scanf("%d %d", &numNums, &cost)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Case #%d: %s\n", n, solve(numNums, cost))
	}
}

func solve(n, c int) string {
	costs := []int{}
	total := 0
	for i := n; i > 1; i-- {
		costs = append(costs, i)
		total += i
	}

	if c < n-1 || c > total {
		return "IMPOSSIBLE"
	}

	for i, cost := range costs {
		if total-c >= cost-1 {
			costs[i] = 1
			total -= cost - 1
		}
	}

	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	for i := n - 2; i >= 0; i-- {
		reverse(nums, i, i+costs[i]-1)
	}

	out := []string{}
	for _, n := range nums {
		out = append(out, strconv.Itoa(n))
	}

	return strings.Join(out, " ")
}

func reverse(nums []int, start, end int) {
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
