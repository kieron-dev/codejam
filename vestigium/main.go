package main

import "fmt"

func main() {
	var numCases int
	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		panic(err)
	}

	for ncase := 0; ncase < numCases; ncase++ {
		var size int
		_, err := fmt.Scanf("%d", &size)
		if err != nil {
			panic(err)
		}

		var rows [][]int
		for r := 0; r < size; r++ {
			var row []int
			for c := 0; c < size; c++ {
				var n int
				_, err := fmt.Scanf("%d", &n)
				if err != nil {
					panic(err)
				}
				row = append(row, n)
			}
			rows = append(rows, row)
		}
		trace, rowRepeats, colRepeats := calc(rows)
		fmt.Printf("Case #%d: %d %d %d\n", ncase+1, trace, rowRepeats, colRepeats)
	}
}

func calc(rows [][]int) (trace, rrep, crep int) {
	for i := 0; i < len(rows); i++ {
		trace += rows[i][i]
	}

	for r := 0; r < len(rows); r++ {
		nums := map[int]bool{}
		for c := 0; c < len(rows); c++ {
			if _, ok := nums[rows[r][c]]; ok {
				rrep++
				break
			}
			nums[rows[r][c]] = true
		}
	}

	for c := 0; c < len(rows); c++ {
		nums := map[int]bool{}
		for r := 0; r < len(rows); r++ {
			if _, ok := nums[rows[r][c]]; ok {
				crep++
				break
			}
			nums[rows[r][c]] = true
		}
	}

	return trace, rrep, crep
}
