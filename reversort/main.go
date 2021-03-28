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
		var numNums int
		_, err := fmt.Scanf("%d", &numNums)
		if err != nil {
			log.Fatal(err)
		}

		nums := []int{}

		for i := 0; i < numNums; i++ {
			var n int
			_, err := fmt.Scanf("%d", &n)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}

		fmt.Printf("Case #%d: %d\n", n, solve(nums))
	}
}

func solve(nums []int) int {
	complexity := 0
	l := len(nums)
	for i := 0; i < l-1; i++ {
		minPos := minPos(nums, i, l-1)
		complexity += minPos - i + 1
		reverse(nums, i, minPos)
	}

	return complexity
}

func minPos(nums []int, start, end int) int {
	min := nums[start]
	minPos := start

	for i := start + 1; i <= end; i++ {
		if nums[i] < min {
			min = nums[i]
			minPos = i
		}
	}

	return minPos
}

func reverse(nums []int, start, end int) {
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
