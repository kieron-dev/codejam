package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var numCases int
	fmt.Scanf("%d\n", &numCases)

	for i := 1; i <= numCases; i++ {
		var d int
		var prog string
		fmt.Scanf("%d %s\n", &d, &prog)
		minSwaps, err := GetMinSwaps(d, prog)
		var answer string
		if err != nil {
			answer = "IMPOSSIBLE"
		} else {
			answer = fmt.Sprintf("%d", minSwaps)
		}
		fmt.Printf("Case #%d: %s\n", i, answer)
	}
}

func GetMinSwaps(d int, prog string) (n int, err error) {
	swaps := 0
	for Score(prog) > d {
		idx := strings.LastIndex(prog, "CS")
		if idx == -1 {
			return 0, errors.New("impossible")
		}
		prog = SwapAtIndex(prog, idx)
		swaps++
	}
	return swaps, nil
}

func SwapAtIndex(prog string, idx int) string {
	return prog[:idx] + "SC" + prog[idx+2:]
}

func Score(prog string) int {
	charge := 1
	damage := 0
	for _, c := range prog {
		if c == 'C' {
			charge *= 2
		}
		if c == 'S' {
			damage += charge
		}
	}
	return damage
}
