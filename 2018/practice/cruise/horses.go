package main

import (
	"fmt"
	"os"
)

func main() {
	var numCases int
	fmt.Fscanf(os.Stdin, "%d\n", &numCases)

	for testcase := 1; testcase <= numCases; testcase++ {
		var d, n int
		fmt.Fscanf(os.Stdin, "%d %d\n", &d, &n)

		var maxTime float64
		for i := 0; i < n; i++ {
			var p, s int
			fmt.Fscanf(os.Stdin, "%d %d\n", &p, &s)
			time := float64(d-p) / float64(s)
			if time > maxTime {
				maxTime = time
			}
		}
		fmt.Printf("Case #%d: %f\n", testcase, float64(d)/maxTime)
	}
}
